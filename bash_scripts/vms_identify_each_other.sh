#!/bin/bash

declare -a allvms=("fa24-cs425-6401.cs.illinois.edu" "fa24-cs425-6402.cs.illinois.edu" "fa24-cs425-6403.cs.illinois.edu" "fa24-cs425-6404.cs.illinois.edu" "fa24-cs425-6405.cs.illinois.edu" "fa24-cs425-6406.cs.illinois.edu" "fa24-cs425-6407.cs.illinois.edu" "fa24-cs425-6408.cs.illinois.edu" "fa24-cs425-6409.cs.illinois.edu" "fa24-cs425-6410.cs.illinois.edu")

for each_vm_machine in "${allvms[@]}"
do
	ssh-copy-id -i ~/.ssh/id_ecdsa.pub "sdevata2@${each_vm_machine}"
done

for each_vm_machine in "${allvms[@]}"
do
	ssh "sdevata2@${each_vm_machine}" -t \"ls\"
done
