#!/bin/bash

declare -a allvms=("fa24-cs425-6401.cs.illinois.edu" "fa24-cs425-6402.cs.illinois.edu" "fa24-cs425-6403.cs.illinois.edu" "fa24-cs425-6404.cs.illinois.edu" "fa24-cs425-6405.cs.illinois.edu" "fa24-cs425-6406.cs.illinois.edu" "fa24-cs425-6407.cs.illinois.edu" "fa24-cs425-6408.cs.illinois.edu" "fa24-cs425-6409.cs.illinois.edu" "fa24-cs425-6410.cs.illinois.edu")

vm_count=1
for each_vm_machine in "${allvms[@]}"
do
   echo "$vm_count $each_vm_machine"
   ls ../data/vm${vm_count}.log
   scp ../data/vm${vm_count}.log "sdevata2@${each_vm_machine}:~/machine.${vm_count}.log"
   scp ../data/machine.${vm_count}.log.test "sdevata2@${each_vm_machine}:~/machine.${vm_count}.log.test"
   (( vm_count++ ))
done
