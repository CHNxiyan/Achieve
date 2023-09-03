# Socratex

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue)](./LICENSE)
[![Node.js Package, Docker Image](https://github.com/Leask/socratex/actions/workflows/build.yml/badge.svg)](https://github.com/Leask/socratex/actions/workflows/build.yml)
[![CodeQL](https://github.com/Leask/socratex/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/Leask/socratex/actions/workflows/codeql-analysis.yml)

A Secure Web Proxy. Which is fast, secure, and easy to use.

*This project is under active development. Everything may change soon.*

<!--img width="1006" alt="Screen Shot 2022-04-15 at 9 40 44 PM" src="https://user-images.githubusercontent.com/233022/163656944-d6b57fbc-8c0d-4cd9-abe0-f65fd39ee728.png"-->

Socratex extends the native [net.createServer](https://nodejs.org/api/net.html#net_net_createserver_options_connectionlistener), and it acts as a real transparent HTTPS-proxy built on top of TCP-level.

It's a real HTTPS proxy, not HTTPS over HTTP. It allows upstream client-request dynamically to other proxies or works as a single layer encrypted proxy.

Socratex will request and set up the certificate automatically, and it will automatically renew the certificate when it expires. You don't need to worry about the dirty work about HTTPS/SSL.

It supports [Basic Proxy-Authentication](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Proxy-Authorization) and Token-Based-Authentication as default. Socratex will create a new token at the first run, you don't need to worry about it.

<img width="952" alt="Screen Shot 2022-04-15 at 8 47 01 PM" src="https://user-images.githubusercontent.com/233022/163655396-720066ba-cb2e-4bd1-9dbc-d167057e5776.png">


## Why another proxy?

First of all, many people in particular countries need proxy software that is easy to deploy and could be used to secure their network traffic. Second, because of the limitation on App Store, especially in China, VPN and proxy software are not allowed to be used. So we need to find a way to avoid censorship without any client apps. Secure Web Proxy is the only choice and a promising one.


## Deploy a Secure Web Proxy within 10 seconds

You need a domain name and set an A-record pointed to your cloud virtual machine.

Usually, that virtual machine can not be located in China.

Assumes that you have a workable Node.js (v16 or above) environment.

Now let's make the magic happen!

- Modern method:
```bash
$ sudo su
# cd ~
# npx socratex --domain=example.com --bypass=cn
```
- Classic method:
```bash
$ git clone git@github.com:Leask/socratex.git
$ cd socratex
$ npm install
$ sudo ./main.mjs --domain=example.com --bypass=cn
```
- With Docker:
```bash
$ touch ~/.socratex.json
$ docker pull leask/socratex
$ docker run -d --restart=always -p 80:80 -p 443:443 \
    -v ~/.socratex.json:/root/.socratex.json \
    leask/socratex --domain=example.com --bypass=cn
```

If everything works fine, you should see a message like this:

```
[SOCRATEX Vx.y.z] https://github.com/Leask/socratex
[SOCRATEX] Secure Web Proxy started at https://example.com:443 (IPv6 ::).
[SOCRATEX] HTTP Server started at http://example.com:80 (IPv6 ::).
[SSL] Creating new private-key and CSR...
[SSL] Done.
[SSL] Updating certificate...
[SSL] Done.
[SOCRATEX] * Token authentication:
[SOCRATEX]   - PAC:  https://example.com/proxy.pac?token=959c298e-9f38-b201-2e7e-14af54469889
[SOCRATEX]   - WPAD: https://example.com/wpad.dat?token=959c298e-9f38-b201-2e7e-14af54469889
[SOCRATEX]   - Log:  https://example.com/console?token=959c298e-9f38-b201-2e7e-14af54469889
[SOCRATEX] * Basic authentication:
[SOCRATEX]   - PAC:   https://foo:bar@example.com/proxy.pac
[SOCRATEX]   - WPAD:  https://foo:bar@example.com/wpad.dat
[SOCRATEX]   - Log:   https://foo:bar@example.com/console
[SOCRATEX]   - Proxy: https://foo:bar@example.com
```

Copy the `PAC url` or `WPAD url` and paste it into your system's `Automatic Proxy Configuration` settings. That is all you need to do.

<img width="780" alt="Screen Shot 2022-04-15 at 5 26 22 PM" src="https://user-images.githubusercontent.com/233022/163655554-a5695c06-97c4-40e9-92d9-dec89bf69805.png">

<img width="948" alt="Screen Shot 2022-04-15 at 5 25 41 PM" src="https://user-images.githubusercontent.com/233022/163655559-1272c3cd-ffe0-4947-b290-326a4862db5d.png">

*Note*: You can also use the `log url` to monitor the system's activity.


## Command line args

All args are optional. In most cases, you just need to set the domain name. Of cause, you can also set the `bypass` countries to reduce proxy traffics.

| Param  | Type                | Description  |
| ------ | ------------------- | ------------ |
| domain | <code>String</code> | Domain to deploy the proxy. |
| http | <code>With/Without</code> | Use HTTP-only-mode for testing only. |
| bypass | <code>String</code> | Bypass IPs in these countries, could be multiple, example: --bypass=CN --bypass=US |
| user | <code>String</code> | Use `user` and `password` to enable Basic Authorization. |
| password | <code>String</code> | Use `user` and `password` to enable Basic Authorization. |
| token | <code>String</code> | Use to enable Token Authorization. |
| address | <code>String</code> | Activate/Handle Proxy-Authentication. Returns or solves to Boolean. |
| port | <code>Number</code> | Default 443 to handle incoming connection. |


## Limitations

### Why not use `sudo npx ...` directly?

Socratex works at default HTTP (80) and HTTPS (443) ports. You need to be root to listen to these ports on some systems. Because of this issue: https://github.com/npm/cli/issues/3110, if you are in a folder NOT OWN by root, you CAN NOT use `sudo npm ...` or `sudo npx ...` directly to run socratex.

### Why doesn't work with iOS?

Socratex can be used with `macOS`, `Chrome OS`, `Windows`, `Linux` and `Android`. But it's NOT compatible with iOS currently. Because iOS does not support `Secure Web Proxy` yet. I will keep an eye on this issue and try any possible walk-around solutions.


## Why name it `Socratex`?

`Socratex` was named after `Socrates,` a Greek philosopher from Athens credited as the founder of Western philosophy and among the first moral philosophers of the ethical tradition of thought.

[![Socrates](https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/David_-_The_Death_of_Socrates.jpg/2880px-David_-_The_Death_of_Socrates.jpg)](https://en.wikipedia.org/wiki/Socrates)

*Image credit: The Death of Socrates, by Jacques-Louis David (1787)*


<details>
    <summary>Programmable proxy</summary>

## Programmable proxy

```
////////////////////////////////////////////////////////////////////////////////
// NO NEED TO READ ANYTHING BELOW IF YOU ARE NOT GOING TO CUSTOMIZE THE PROXY //
////////////////////////////////////////////////////////////////////////////////
```

You can also use socratex as a programmable proxy to meet your own needs.

```bash
$ npm i -s socratex
```

Socratex is an ES6 module, so you can use it in your modern Node.js projects.

```javascript
import { Socratex } from 'socratex';

const [port, address, options] = ['4698', '': {}];

const socratex = new Socratex(options);

socratex.listen(port, address, async () => {
    console.log('TCP-Proxy-Server started at: ', server.address());
});
```


### Options object use to customize the proxy

`options` should be an object.

| Param  | Type                | Description  |
| ------ | ------------------- | ------------ |
| basicAuth | <code>Function/AsyncFunction</code> | Activate/Handle Proxy-Authentication. Returns or solves to Boolean. |
| tokenAuth | <code>Function/AsyncFunction</code> | Activate/Handle Proxy-Authentication. Returns or solves to Boolean. |
| upstream | <code>Function/AsyncFunction</code> | The proxy to be used to upstreaming requests. Returns String. |
| tcpOutgoingAddress | <code>Function/AsyncFunction</code> | The localAddress to use while sending requests. Returns String. |
| injectData | <code>Function/AsyncFunction</code> | The edited data to upstream. Returns Buffer or string. |
| injectResponse | <code>Function/AsyncFunction</code> | The edited response to return to connected client. Returns Buffer or string. |
| keys | <code>Function/AsyncFunction</code> | The keys to use while handshake. It will work only if intercept is true. Returns Object or false. |
| logLevel | <code>Number</code> | Default 0 to log all messages. |
| intercept | <code>Boolean</code> | Activate interception of encrypted communications. False as default. |


### `upstream`, `tcpOutgoingAddress`, `injectData` & `injectResponse` options

The options are functions having follow parameters:

| Param  | Type                | Description  |
| ------ | ------------------- | ------------ |
| data | <code>Buffer</code> | The received data. |
| session | <code>Session</code> | Object containing info/data about Tunnel. |

- upstream-Function need to return/resolve a String with format -> `IP:PORT` or `USER:PWD@IP:PORT` of used http-proxy. If *'localhost'* is returned/resolved, then the host-self will be used as proxy.
- tcpOutgoingAddress-Function need to return a String with format -> `IP`.
- injectData-Function need to return a String or buffer for the new spoofed data. This will be upstreamed as request.
- injectResponse-Function need to return a String or buffer for the new received data.

*Note*: These functions will be executed before first tcp-socket-connection is established.


### Upstream to other proxies

If you don't want to use the host of active instance self, then you need to upstream connections to another http-proxy.
This can be done with `upstream` attribute.

```javascript
const options = {
    upstream: async () => { return 'x.x.x.x:3128'; },
};
```

### The Basic Authorization mechanism

This activate basic authorization mechanism.
The Auth-function will be executed while handling Proxy-Authentications.

| Param  | Type                | Description  |
| ------ | ------------------- | ------------ |
|username | <code>String</code> |  The client username. |
|password | <code>String</code> |  The client password |
|session | <code>Session</code> |  Object containing info/data about Tunnel |

*Note*: It needs to return True/False or a **Promise** that resolves to boolean (*isAuthenticated*).

```javascript
const options = {
    basicAuth: async (user, password) => user === 'bar' && password === 'foo';
};
```

### The Token Authorization mechanism

This activate token authorization mechanism.
The Auth-function will be executed while handling Proxy-Authentications.

| Param  | Type                | Description  |
| ------ | ------------------- | ------------ |
| token | <code>String</code> |  The client token. |
| session | <code>Session</code> |  Object containing info/data about Tunnel |

*Note*: It needs to return True/False or a **Promise** that resolves to boolean (*isAuthenticated*).

```javascript
const options = {
    tokenAuth: async (token) => token === 'a-very-long-token';
};
```

### Interception

This feature is in very early stage, and it's for web development only. The callbacks `injectData` & `injectResponse` could be used to intercept/spoof communication. These functions are executed with the `data` and `session` arguments.

### Intercepting HTTPS

The boolean attribute `intercept` allows to break SSL-Communication between Source & Destination. This will activate Security-Alarm by most used browsers.

```javascript
const [uaToSwitch, switchWith] = ['curl 7.79.1', 'a-fake-user-agent'];
const options = {
    intercept: true,
    injectData(data, session) {
        if (session.isHttps && data.toString().match(uaToSwitch)) {
            return Buffer.from(data.toString().replace(uaToSwitch, switchWith));
        }
        return data;
    },
};
```

```bash
curl -x localhost:8080 -k http://ifconfig.io/ua
curl 7.79.1

curl -x localhost:8080 -k https://ifconfig.me/ua
a-fake-user-agent
```


### The `keys` Function

You can use this option to provide your own self-signed certificate.

If activated needs to return an Object `{key:'String', cert:'String'}` like [native tls_connect_options.key & tls_connect_options.cert](https://nodejs.org/api/tls.html#tls_tls_connect_options_callback) or `false` statement.

If no object is returned, then [default keys](https://github.com/Leask/socratex/tree/main/keys) will be used to update communication.

| Param  | Type                | Description  |
| ------ | ------------------- | ------------ |
| session | <code>Session</code> | Object containing info/data about Tunnel. |

*Note*: This function will be executed before TLS-Handshake.

### Session-instance

The Session-instance is a Object containing info/data about Tunnel.

Use `.getConnections()` to get the current connections.

```javascript
setInterval(() => {
    const connections = socratex.getConnections();
    console.log([new Date()], 'OPEN =>', Object.keys(connections).length)
}, 3000);
```

The connection items in the connections array include useful attributes/methods:

- isHttps - Is session encrypted.
- getTunnelStats() - Get Stats for this tunnel
- getId() - Get Own ID-Session
- isAuthenticated() - Is the session authenticated by user or not.
- ... (More APIS tobe documented)


### Dynamically routing

This example upstreams only requests for ifconfig.me to another proxy, for all other requests will be used localhost.

```javascript
const options = {
    upstream(data, session) {
        return data.toString().includes('ifconfig.me')
            ? 'x.x.x.x:3128' : 'localhost';
    },
});
```

Testing with `curl`:

```bash
curl -x 127.0.0.1:8080 https://ifconfig.me
x.x.x.x

curl -x 127.0.0.1:8080 https://ifconfig.co
y.y.y.y
```

</details>
