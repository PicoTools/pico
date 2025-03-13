# Pico C2

![Pico CLI](.github/assets/pico.png)

Pico is a C2 framework that allows you to connect your agents and listeners written in any programming language.

> :warning: Pico C2 is in the early stages of development. The API and functionality may change over time!

## Why Another C2 Framework?

Unlike other C2 frameworks, such as Cobalt Strike, Pico does not offer pre-built listeners or agents. Instead, you have full control over the development process. You can create your own custom listener and payloads, implement a communication protocol that suits your needs, and then integrate these components with Pico using gRPC.

Also, Pico includes its own programming language, [PLAN](https://github.com/PicoTools/plan), which is designed for automating tasks (similar to agressor scripts for Cobalt Strike).

## Quick Start

1. Download **Pico** from [latest release](https://github.com/PicoTools/pico/releases) or build it yourself (`make build`).
2. Create a server configuration file [(Example)](https://github.com/PicoTools/pico/blob/master/config/config.yml).
3. Start the server:

```sh
$ ./pico --config config.yml run
000.01 I listener start serving {"ip": "0.0.0.0", "port": 51235, "fingerprint": "ecc67e3a0b6db4ecee4ff6aa195fec0719a62bdc"}
000.01 I operator start serving {"ip": "0.0.0.0", "port": 51234, "fingerprint": "db4505125dbdb8c84667a17ecd7c50dda7dfb6f6"}
000.01 I management start serving {"ip": "0.0.0.0", "port": 51233, "fingerprint": "b21620d6457d107b06e03ab9128cead5ae61aacd", "token": "bo4wro8CEvWzlsJtQgrHdkgIgw9lIPd0"}
...
```

4. Download **Pico-ctl** from [latest release](https://github.com/PicoTools/pico/releases) or build it yourself (`make build-ctl`).
5. Use **Pico-ctl** to generate tokens for the operator and listener. You can find the management token in the server logs.

```sh
$ ./pico-ctl -H 127.0.0.1:51233 -t bo4wro8CEvWzlsJtQgrHdkgIgw9lIPd0
[pico-ctl] > operator add pico-operator
Username:  pico-operator
Token:     Y24IQl9Wp2COnV8lIDHyucIuX2Z5VJlP
Last:      [never]

[pico-ctl] > listener add
ID:        2
Token:     DrbqGN7Ulfu0qpYMpLDR74BemE2WtXFP
Name:      [none]
IP:        [none]
Port:      [none]
Last:      [none]
```

6. You can use [Pico-cli](https://github.com/PicoTools/pico-cli) to connect the operator to the server.

```sh
$ ./pico-cli -H 127.0.0.1:51234 -t Y24IQl9Wp2COnV8lIDHyucIuX2Z5VJlP
pico > whoami
pico-operator
```

## Agent & Listener Development

Check source code of [example agent](https://github.com/PicoTools/example-agent) and [example listener](https://github.com/PicoTools/example-listener)

*TODO...*

## Contributing

*TODO...*

## Note

**This software is developed for educational and research purposes only.**

*The author does not endorse, promote, or condone any illegal activities. By using this tool, you agree that you are solely responsible for any actions performed with it. The author assumes no liability for any misuse, damages, or legal consequences resulting from the use of this software. Use this tool only in environments where you have explicit permission, such as security research, penetration testing, or red team operations authorized by the target organization. Ensure compliance with all applicable laws and regulations.*
