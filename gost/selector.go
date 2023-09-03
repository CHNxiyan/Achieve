package gost

import (
	"errors"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// ErrNoneAvailable indicates there is no node available.
	ErrNoneAvailable = errors.New("none node available")
)

// NodeSelector as a mechanism to pick nodes and mark their status.
type NodeSelector interface {
	Select(nodes []Node, opts ...SelectOption) (Node, error)
}

type defaultSelector struct {
}

func (s *defaultSelector) Select(nodes []Node, opts ...SelectOption) (Node, error) {
	sopts := SelectOptions{}
	for _, opt := range opts {
		opt(&sopts)
	}

	for _, filter := range sopts.Filters {
		nodes = filter.Filter(nodes)
	}
	if len(nodes) == 0 {
		return Node{}, ErrNoneAvailable
	}
	strategy := sopts.Strategy
	if strategy == nil {
		strategy = &RoundStrategy{}
	}
	return strategy.Apply(nodes), nil
}

// SelectOption is the option used when making a select call.
type SelectOption func(*SelectOptions)

// SelectOptions is the options for node selection.
type SelectOptions struct {
	Filters  []Filter
	Strategy Strategy
}

// WithFilter adds a filter function to the list of filters
// used during the Select call.
func WithFilter(f ...Filter) SelectOption {
	return func(o *SelectOptions) {
		o.Filters = append(o.Filters, f...)
	}
}

// WithStrategy sets the selector strategy
func WithStrategy(s Strategy) SelectOption {
	return func(o *SelectOptions) {
		o.Strategy = s
	}
}

// Strategy is a selection strategy e.g random, round-robin.
type Strategy interface {
	Apply([]Node) Node
	String() string
}

// NewStrategy creates a Strategy by the name s.
func NewStrategy(s string) Strategy {
	switch s {
	case "random":
		return &RandomStrategy{}
	case "fifo":
		return &FIFOStrategy{}
	case "round":
		fallthrough
	default:
		return &RoundStrategy{}
	}
}

// RoundStrategy is a strategy for node selector.
// The node will be selected by round-robin algorithm.
type RoundStrategy struct {
	counter uint64
}

// Apply applies the round-robin strategy for the nodes.
func (s *RoundStrategy) Apply(nodes []Node) Node {
	if len(nodes) == 0 {
		return Node{}
	}

	n := atomic.AddUint64(&s.counter, 1) - 1
	return nodes[int(n%uint64(len(nodes)))]
}

func (s *RoundStrategy) String() string {
	return "round"
}

// RandomStrategy is a strategy for node selector.
// The node will be selected randomly.
type RandomStrategy struct {
	Seed int64
	rand *rand.Rand
	once sync.Once
	mux  sync.Mutex
}

// Apply applies the random strategy for the nodes.
func (s *RandomStrategy) Apply(nodes []Node) Node {
	s.once.Do(func() {
		seed := s.Seed
		if seed == 0 {
			seed = time.Now().UnixNano()
		}
		s.rand = rand.New(rand.NewSource(seed))
	})
	if len(nodes) == 0 {
		return Node{}
	}

	s.mux.Lock()
	r := s.rand.Int()
	s.mux.Unlock()

	return nodes[r%len(nodes)]
}

func (s *RandomStrategy) String() string {
	return "random"
}

// FIFOStrategy is a strategy for node selector.
// The node will be selected from first to last,
// and will stick to the selected node until it is failed.
type FIFOStrategy struct{}

// Apply applies the fifo strategy for the nodes.
func (s *FIFOStrategy) Apply(nodes []Node) Node {
	if len(nodes) == 0 {
		return Node{}
	}
	return nodes[0]
}

func (s *FIFOStrategy) String() string {
	return "fifo"
}

// Filter is used to filter a node during the selection process
type Filter interface {
	Filter([]Node) []Node
	String() string
}

// default options for FailFilter
const (
	DefaultMaxFails    = 1
	DefaultFailTimeout = 30 * time.Second
)

// FailFilter filters the dead node.
// A node is marked as dead if its failed count is greater than MaxFails.
type FailFilter struct {
	MaxFails    int
	FailTimeout time.Duration
}

// Filter filters dead nodes.
func (f *FailFilter) Filter(nodes []Node) []Node {
	maxFails := f.MaxFails
	if maxFails == 0 {
		maxFails = DefaultMaxFails
	}
	failTimeout := f.FailTimeout
	if failTimeout == 0 {
		failTimeout = DefaultFailTimeout
	}

	if len(nodes) <= 1 || maxFails < 0 {
		return nodes
	}
	nl := []Node{}
	for i := range nodes {
		marker := nodes[i].marker.Clone()
		// log.Logf("%s: %d/%d %v/%v", nodes[i], marker.FailCount(), f.MaxFails, marker.FailTime(), f.FailTimeout)
		if marker.FailCount() < uint32(maxFails) ||
			time.Since(time.Unix(marker.FailTime(), 0)) >= failTimeout {
			nl = append(nl, nodes[i])
		}
	}
	return nl
}

func (f *FailFilter) String() string {
	return "fail"
}

// InvalidFilter filters the invalid node.
// A node is invalid if its port is invalid (negative or zero value).
type InvalidFilter struct{}

// Filter filters invalid nodes.
func (f *InvalidFilter) Filter(nodes []Node) []Node {
	nl := []Node{}
	for i := range nodes {
		_, sport, _ := net.SplitHostPort(nodes[i].Addr)
		if port, _ := strconv.Atoi(sport); port > 0 {
			nl = append(nl, nodes[i])
		}
	}
	return nl
}

func (f *InvalidFilter) String() string {
	return "invalid"
}

type failMarker struct {
	failTime  int64
	failCount uint32
	mux       sync.RWMutex
}

func (m *failMarker) FailTime() int64 {
	if m == nil {
		return 0
	}

	m.mux.Lock()
	defer m.mux.Unlock()

	return m.failTime
}

func (m *failMarker) FailCount() uint32 {
	if m == nil {
		return 0
	}

	m.mux.Lock()
	defer m.mux.Unlock()

	return m.failCount
}

func (m *failMarker) Mark() {
	if m == nil {
		return
	}

	m.mux.Lock()
	defer m.mux.Unlock()

	m.failTime = time.Now().Unix()
	m.failCount++
}

func (m *failMarker) Reset() {
	if m == nil {
		return
	}

	m.mux.Lock()
	defer m.mux.Unlock()

	m.failTime = 0
	m.failCount = 0
}

func (m *failMarker) Clone() *failMarker {
	if m == nil {
		return nil
	}

	m.mux.RLock()
	defer m.mux.RUnlock()

	fc, ft := m.failCount, m.failTime

	return &failMarker{
		failCount: fc,
		failTime:  ft,
	}
}
