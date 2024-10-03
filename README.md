# Distributed Grep
This repository contains an implementation of Distributed Grep over a cluster of VMs. A search query on any one VM, efficiently queries all connected VMs and returns an aggregated response.

To run, in the home folder execute

```
go run .
```

Next, in the prompt simply enter the regex query you would like to search for.

For example:
```
Enter grep regex pattern (in single/double quotes):"DEBUG"
```

## Running Tests
In the home directory, execute
```
go test
```

## Setting up the VMs
There are helper scripts to set up the VMs.

### Copying logs to respective VMs
```
cd bash_scripts; bash copy_logs_to_vms.sh
```

### Copying code folder to respective VMs
```
cd bash_scripts; bash copy_to_vms.sh
```

### Adding SSH keys of the VMs to each other (run on each VM)
```
cd bash_scripts; bash vms_identify_each_other.sh
```

