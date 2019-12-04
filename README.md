Reach
=====

```bash
reach [-d <file_path|url>] <script.tg> [arg1] [arg2] ....
```

```tengo
// script.tg
for c in Data["edi_app_servers"] {
    Go(func() {
        h := Connect(c)
        h.Log('installing system libraries')  // prefix with time, host, module?
        h.Do('apt', {present: True, name: 'libpq-dev'}, {ignore_failure}) // customize running modules
        h.Do('template', {template: 'some.tpl', data: {}, dest: '/opt/edi/env.sh'})
    })
}
```

```tengo
// apt.tg
func(h, data) {
    if data['present'] {
        h.Do('shell', {command: 'apt-get install -y ' + data['name']})
    } else {
        h.Do('shell', {command: 'apt-get remove -y ' + data['name']})
    }
}
```

Global Primitives:
- Go
- Register
- Host API:
  - Do
  - Log

Main script Primitives:
- Args
- Data
- Connect

Good to have:
- Compile modules before startup to validate modules are ok
- "Prerun" the modules to track module dependencies and validate all dependencies are met
- Detailed logging from connection success/failures and module use from lowest to highest (with results)
- Module stack trace?
