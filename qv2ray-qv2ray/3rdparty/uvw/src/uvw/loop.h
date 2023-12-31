#ifndef UVW_LOOP_INCLUDE_H
#define UVW_LOOP_INCLUDE_H


#ifdef _WIN32
#include <ciso646>
#endif

#include <functional>
#include <memory>
#include <utility>
#include <type_traits>
#include <chrono>
#include <uv.h>
#include "emitter.h"
#include "util.h"


namespace uvw {


class AsyncHandle;
class CheckHandle;
class FsEventHandle;
class FsPollHandle;
class IdleHandle;
class PipeHandle;
class PollHandle;
class PrepareHandle;
class ProcessHandle;
class SignalHandle;
class TCPHandle;
class TimerHandle;
class TTYHandle;
class UDPHandle;


namespace details {


enum class UVLoopOption: std::underlying_type_t<uv_loop_option> {
    BLOCK_SIGNAL = UV_LOOP_BLOCK_SIGNAL
};


enum class UVRunMode: std::underlying_type_t<uv_run_mode> {
    DEFAULT = UV_RUN_DEFAULT,
    ONCE = UV_RUN_ONCE,
    NOWAIT = UV_RUN_NOWAIT
};


}


/**
 * @brief The Loop class.
 *
 * The event loop is the central part of `uvw`'s functionalities, as well as
 * `libuv`'s ones.<br/>
 * It takes care of polling for I/O and scheduling callbacks to be run based on
 * different sources of events.
 */
class Loop final: public Emitter<Loop>, public std::enable_shared_from_this<Loop> {
    using Deleter = void(*)(uv_loop_t *);

    template<typename, typename>
    friend class Resource;

    template<typename R, typename... Args>
    auto create_resource(int, Args&&... args) -> decltype(std::declval<R>().init(), std::shared_ptr<R>{}) {
        auto ptr = R::create(shared_from_this(), std::forward<Args>(args)...);
        ptr = ptr->init() ? ptr : nullptr;
        return ptr;
    }

    template<typename R, typename... Args>
    std::shared_ptr<R> create_resource(char, Args&&... args) {
        return R::create(shared_from_this(), std::forward<Args>(args)...);
    }

    Loop(std::unique_ptr<uv_loop_t, Deleter> ptr) noexcept;

public:
    using Time = std::chrono::duration<uint64_t, std::milli>;
    using Configure = details::UVLoopOption;
    using Mode = details::UVRunMode;

    /**
     * @brief Initializes a new Loop instance.
     * @return A pointer to the newly created loop.
     */
    static std::shared_ptr<Loop> create();

    /**
     * @brief Initializes a new Loop instance from an existing resource.
     *
     * The lifetime of the resource must exceed that of the instance to which
     * it's associated. Management of the memory associated with the resource is
     * in charge of the user.
     *
     * @param loop A valid pointer to a correctly initialized resource.
     * @return A pointer to the newly created loop.
     */
    static std::shared_ptr<Loop> create(uv_loop_t *loop);

    /**
     * @brief Gets the initialized default loop.
     *
     * It may return an empty pointer in case of failure.<br>
     * This function is just a convenient way for having a global loop
     * throughout an application, the default loop is in no way different than
     * the ones initialized with `create()`.<br>
     * As such, the default loop can be closed with `close()` so the resources
     * associated with it are freed (even if it is not strictly necessary).
     *
     * @return The initialized default loop.
     */
    static std::shared_ptr<Loop> getDefault();

    Loop(const Loop &) = delete;
    Loop(Loop &&other) = delete;
    Loop & operator=(const Loop &) = delete;
    Loop & operator=(Loop &&other) = delete;

    ~Loop() noexcept;

    /**
     * @brief Sets additional loop options.
     *
     * You should normally call this before the first call to uv_run() unless
     * mentioned otherwise.<br/>
     * Supported options:
     *
     * * `Loop::Configure::BLOCK_SIGNAL`: Block a signal when polling for new
     * events. A second argument is required and it is the signal number.
     *
     * An ErrorEvent will be emitted in case of errors.
     *
     * See the official
     * [documentation](http://docs.libuv.org/en/v1.x/loop.html#c.uv_loop_configure)
     * for further details.
     */
    template<typename... Args>
    void configure(Configure flag, Args&&... args) {
        auto option = static_cast<std::underlying_type_t<Configure>>(flag);
        auto err = uv_loop_configure(loop.get(), static_cast<uv_loop_option>(option), std::forward<Args>(args)...);
        if(err) { publish(ErrorEvent{err}); }
    }

    /**
     * @brief Creates resources of any type.
     *
     * This should be used as a default method to create resources.<br/>
     * The arguments are the ones required for the specific resource.
     *
     * Use it as `loop->resource<uvw::TimerHandle>()`.
     *
     * @return A pointer to the newly created resource.
     */
    template<typename R, typename... Args>
    std::shared_ptr<R> resource(Args&&... args) {
        return create_resource<R>(0, std::forward<Args>(args)...);
    }

    /**
     * @brief Releases all internal loop resources.
     *
     * Call this function only when the loop has finished executing and all open
     * handles and requests have been closed, or the loop will emit an error.
     *
     * An ErrorEvent will be emitted in case of errors.
     */
    void close();

    /**
     * @brief Runs the event loop.
     *
     * Available modes are:
     *
     * * `Loop::Mode::DEFAULT`: Runs the event loop until there are no more
     * active and referenced handles or requests.
     * * `Loop::Mode::ONCE`: Poll for i/o once. Note that this function blocks
     * if there are no pending callbacks.
     * * `Loop::Mode::NOWAIT`: Poll for i/o once but don’t block if there are no
     * pending callbacks.
     *
     * See the official
     * [documentation](http://docs.libuv.org/en/v1.x/loop.html#c.uv_run)
     * for further details.
     *
     * @return True when done, false in all other cases.
     */
    template<Mode mode = Mode::DEFAULT>
    bool run() noexcept;

    /**
     * @brief Checks if there are active resources.
     * @return True if there are active resources in the loop.
     */
    bool alive() const noexcept;

    /**
     * @brief Stops the event loop.
     *
     * It causes `run()` to end as soon as possible.<br/>
     * This will happen not sooner than the next loop iteration.<br/>
     * If this function was called before blocking for I/O, the loop won’t block
     * for I/O on this iteration.
     */
    void stop() noexcept;

    /**
     * @brief Get backend file descriptor.
     *
     * Only kqueue, epoll and event ports are supported.<br/>
     * This can be used in conjunction with `run<Loop::Mode::NOWAIT>()` to poll
     * in one thread and run the event loop’s callbacks in another.
     *
     * @return The backend file descriptor.
     */
    int descriptor() const noexcept;

    /**
     * @brief Gets the poll timeout.
     * @return A `std::pair` composed as it follows:
     * * A boolean value that is true in case of valid timeout, false otherwise.
     * * Milliseconds (`std::chrono::duration<uint64_t, std::milli>`).
     */
    std::pair<bool, Time> timeout() const noexcept;

    /**
     * @brief Returns the current timestamp in milliseconds.
     *
     * The timestamp is cached at the start of the event loop tick.<br/>
     * The timestamp increases monotonically from some arbitrary point in
     * time.<br/>
     * Don’t make assumptions about the starting point, you will only get
     * disappointed.
     *
     * @return The current timestamp in milliseconds (actual type is
     * `std::chrono::duration<uint64_t, std::milli>`).
     */
    Time now() const noexcept;

    /**
     * @brief Updates the event loop’s concept of _now_.
     *
     * The current time is cached at the start of the event loop tick in order
     * to reduce the number of time-related system calls.<br/>
     * You won’t normally need to call this function unless you have callbacks
     * that block the event loop for longer periods of time, where _longer_ is
     * somewhat subjective but probably on the order of a millisecond or more.
     */
    void update() const noexcept;

    /**
     * @brief Walks the list of handles.
     *
     * The callback will be executed once for each handle that is still active.
     *
     * @param callback A function to be invoked once for each active handle.
     */
    template<typename Func>
    void walk(Func callback) {
        // remember: non-capturing lambdas decay to pointers to functions
        uv_walk(loop.get(), [](uv_handle_t *handle, void *func) {
            auto &cb = *static_cast<Func *>(func);

            switch(Utilities::guessHandle(HandleCategory{handle->type})) {
            case HandleType::ASYNC:
                cb(*static_cast<AsyncHandle *>(handle->data));
                break;
            case HandleType::CHECK:
                cb(*static_cast<CheckHandle *>(handle->data));
                break;
            case HandleType::FS_EVENT:
                cb(*static_cast<FsEventHandle *>(handle->data));
                break;
            case HandleType::FS_POLL:
                cb(*static_cast<FsPollHandle *>(handle->data));
                break;
            case HandleType::IDLE:
                cb(*static_cast<IdleHandle *>(handle->data));
                break;
            case HandleType::PIPE:
                cb(*static_cast<PipeHandle *>(handle->data));
                break;
            case HandleType::POLL:
                cb(*static_cast<PollHandle *>(handle->data));
                break;
            case HandleType::PREPARE:
                cb(*static_cast<PrepareHandle *>(handle->data));
                break;
            case HandleType::PROCESS:
                cb(*static_cast<ProcessHandle *>(handle->data));
                break;
            case HandleType::SIGNAL:
                cb(*static_cast<SignalHandle *>(handle->data));
                break;
            case HandleType::TCP:
                cb(*static_cast<TCPHandle *>(handle->data));
                break;
            case HandleType::TIMER:
                cb(*static_cast<TimerHandle *>(handle->data));
                break;
            case HandleType::TTY:
                cb(*static_cast<TTYHandle *>(handle->data));
                break;
            case HandleType::UDP:
                cb(*static_cast<UDPHandle *>(handle->data));
                break;
            default:
                // returns the underlying handle, uvw doesn't manage it properly yet
                cb(handle);
                break;
            }
        }, &callback);
    }

    /**
     * @brief Reinitialize any kernel state necessary in the child process after
     * a fork(2) system call.
     *
     * Previously started watchers will continue to be started in the child
     * process.
     *
     * It is necessary to explicitly call this function on every event loop
     * created in the parent process that you plan to continue to use in the
     * child, including the default loop (even if you don’t continue to use it
     * in the parent). This function must be called before calling any API
     * function using the loop in the child. Failure to do so will result in
     * undefined behaviour, possibly including duplicate events delivered to
     * both parent and child or aborting the child process.
     *
     * When possible, it is preferred to create a new loop in the child process
     * instead of reusing a loop created in the parent. New loops created in the
     * child process after the fork should not use this function.
     *
     * Note that this function is not implemented on Windows.<br/>
     * Note also that this function is experimental in `libuv`. It may contain
     * bugs, and is subject to change or removal. API and ABI stability is not
     * guaranteed.
     *
     * An ErrorEvent will be emitted in case of errors.
     *
     * See the official
     * [documentation](http://docs.libuv.org/en/v1.x/loop.html#c.uv_loop_fork)
     * for further details.
     */
    void fork() noexcept;

    /**
     * @brief Gets user-defined data. `uvw` won't use this field in any case.
     * @return User-defined data if any, an invalid pointer otherwise.
     */
    template<typename R = void>
    std::shared_ptr<R> data() const {
        return std::static_pointer_cast<R>(userData);
    }

    /**
     * @brief Sets arbitrary data. `uvw` won't use this field in any case.
     * @param uData User-defined arbitrary data.
     */
    void data(std::shared_ptr<void> uData);

    /**
     * @brief Gets the underlying raw data structure.
     *
     * This function should not be used, unless you know exactly what you are
     * doing and what are the risks.<br/>
     * Going raw is dangerous, mainly because the lifetime management of a loop,
     * a handle or a request is in charge to the library itself and users should
     * not work around it.
     *
     * @warning
     * Use this function at your own risk, but do not expect any support in case
     * of bugs.
     *
     * @return The underlying raw data structure.
     */
    const uv_loop_t * raw() const noexcept;

    /**
     * @brief Gets the underlying raw data structure.
     *
     * This function should not be used, unless you know exactly what you are
     * doing and what are the risks.<br/>
     * Going raw is dangerous, mainly because the lifetime management of a loop,
     * a handle or a request is in charge to the library itself and users should
     * not work around it.
     *
     * @warning
     * Use this function at your own risk, but do not expect any support in case
     * of bugs.
     *
     * @return The underlying raw data structure.
     */
    uv_loop_t * raw() noexcept;

private:
    std::unique_ptr<uv_loop_t, Deleter> loop;
    std::shared_ptr<void> userData{nullptr};
};


// (extern) explicit instantiations

extern template bool Loop::run<Loop::Mode::DEFAULT>() noexcept;
extern template bool Loop::run<Loop::Mode::ONCE>() noexcept;
extern template bool Loop::run<Loop::Mode::NOWAIT>() noexcept;


}


#ifndef UVW_AS_LIB
#include "loop.cpp"
#endif

#endif // UVW_LOOP_INCLUDE_H
