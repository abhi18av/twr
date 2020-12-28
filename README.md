# Installation

Simply download and unzip the release archive for your platform.

NOTE: You need to execute `sudo spctl --master-disable` to allow the binary to be executed on modern macOS.

# USAGE

```
$ twr_0.0.9_darwin_amd64 eklavya$ ./twr
The twr CLI can be configured to be used with a Nextflow Tower installation.
By default, it is configured to be used with the public tower.nf instance.

Usage:
  twr [command]

Available Commands:
  help        Help about any command
  serviceInfo Brief information for the Tower installation.

Flags:
      --config string   config file (default is $HOME/.twr.yaml)
  -h, --help            help for twr

Use "twr [command] --help" for more information about a command.

```


## Config file

Create a `.twr.yaml` file in your $HOME with the following content

```yaml
TOWER_API_ENDPOINT:
  https://api.tower.nf/
```

## Commands

- `serviceInfo`

This command can be used to obtain information about a Tower installation.

```
$ twr_0.0.9_darwin_amd64 eklavya$ ./twr serviceInfo
Using config file: /Users/eklavya/.twr.yaml
{
  "serviceInfo": {
    "version": "20.12.1",
    "commitId": "6b2b93a",
    "authTypes": [
      "github",
      "google"
    ],
    "loginPath": "/login",
    "navbar": {
      "menus": [
        {
          "label": "Docs",
          "url": "https://help.tower.nf"
        },
        {
          "label": "Community",
          "url": "https://gitter.im/nf-tower/community"
        },
        {
          "label": "Feedback",
          "url": "https://github.com/seqeralabs/nf-tower/issues"
        },
        {
          "label": "Support",
          "url": "https://seqera.io/"
        }
      ]
    },
    "heartbeatInterval": 3240
  }
}
```

## TODO

- `actions`
- `computerEnvs`
- `credentials`
- `ga4gh`
- `launch`
- `platforms`
- `tokens`
- `trace`
- `users`
- `workflow`

  
