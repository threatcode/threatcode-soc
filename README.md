# Threat Code Console

Threat Code Console is a customized fork of Threatsensor. This product provides the new primary user interface for the Threat Code product suite in addition to managing the environment's security sensors, as was the original intention of Threatsensor. The following information is from the original Threatsensor readme which is currently still applicable:

# Threatsensor

This system is comprised of a server and one or more agents. An agent runs on a remote server or device and performs jobs that have been queued on the server. The server provides a user interface for administrators to perform various tasks, such as:

* monitor all sensors
* add new jobs
* view the job queue
* inspect job details

## License

Threatsensor is distributed under the terms of version 2 of the GNU General Public License (GPL).

Only version 2 of the GNU GPL license applies to Threatsensor. No other GNU GPL version applies.

See [LICENSE](https://raw.githubusercontent.com/threatcode/threatcode-soc/master/LICENSE) for more information.

## Design

Threatsensor, at its core, is a job queueing and dispatch system. It is targetted specifically towards the computer security industry, primarily due to the job's focus on network concepts. For example, when creating a job, the source address and port, destination address and port, and beginning and ending time range can be specified. But additional, untyped job parameters can also be specified, which allows a job to target systems outside of the security realm.

Each job must be assigned to a sensor. Sensors are responsible for periodically checking into the server to let the server know that the sensor is still up, and to check for new jobs that it must process. Upon completion of a job, the sensor will update the server with the job results.

Job results are provided in any form, and are encapsulated within a byte stream. The job results typically represent a packet structure, in PCAP form.

## Modules

In order to allow Threatsensor to support the variety of job types that are expected to be used, specific functionality of certain job types are isolated to optional modules. However, some module types are mandatory. For example, all servers and agents must provide an auth module, while the server must also provide a datastore module. This doesn't mean a specific module must be enabled, but rather that a module that implements a particular interface must be enabled. This provides runtime flexibility in how each Threatsensor deployment performs.

Threatsensor's configuration file controls whether additional modules should be enabled. Modules can provide custom, or specific integration with other systems. Both the server and the agent have their own lists of available modules.

Below are the current modules:

### filedatastore (server)

Implements the Datastore module by using the local file system for persisting information about pending and completed jobs.

### statickeyauth (server and agent)

Implements the Auth module by specifying a shared key that both the agents and the server will utilize, as well as an optional CIDR block that, if specified, allows the server to bypass the key verification if the request originates from a network in the CIDR block.

### elastic (server)

A custom module that integrates Threatsensor with [Elasticsearch](https://elastic.co). This module provides a new endpoint, /joblookup?esid=xyz, where xyz is the Elasticsearch document ID containing the desired packet source, target, and date ranges to retrieve.

### stenoquery (agent)

Integrates with [Google Stenographer](https://github.com/google/stenographer) for retrieving PCAP files from remote sensors.

### kratos

Manage users stored within the Kratos identity management server. Currently supports listing and deleting users. Updates to users is still in progress. The kratos module requires a hostUrl config parameter, which represents the Kratos admin URL.

## Reverse Proxy

While the integration to an identity server is in progress, the current state of Threatsensor continues to require a reverse proxy, such as [NGINX](https://www.nginx.com/), be deployed in front of the Threatsensor server. 

## Configuration

A single compiled Threatsensor binary can be configured to run either as a server, an agent, or as both (for development and test purposes.) See the example configuration file below for an example configuration of the dual agent and server configuration, and with an additional SecurityOnion module.

Sample threatsensor.json file:
```json
{
  "logLevel": "debug",
  "logFilename": "threatsensor.log",
  "server": {
    "bindAddress": "192.168.1.10:9822",
    "baseUrl": "/threatsensor",
    "maxPacketCount": 5000,
    "htmlDir": "html",
    "modules": {
      "filedatastore": {
        "jobDir": "jobs"
      },
      "elastic": {
        "hostUrl": "http://10.66.166.100:9200",
        "username": "",
        "password": "",
        "verifyCert": false
      },
      "statickeyauth": {
        "anonymousCidr": "192.168.2.103/0",
        "apiKey": "123abc"
      }
    }
  },
  "agent": {
    "pollIntervalMs": 10000,
    "serverUrl": "http://192.168.1.10:9822",
    "modules": {
      "statickeyauth": {
        "apiKey": "123abc"
      },
      "stenoquery": {
        "pcapInputPath": "/nsm/pcap",
        "pcapOutputPath": "/nsm/pcapoutput"
      }
    }
  }
}
```

In a real deployment, the server block and agent block would not coexist in the same configuration file.

The agent can be configured to use a proxy by ensuring the HTTP_PROXY environment variable is provided. Ex:

```bash
export HTTP_PROXY="http://proxyIp:proxyPort"
```
