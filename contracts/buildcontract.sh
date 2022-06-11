#!/bin/sh

#dir_compilation=./compilation
#if [ ! -d "$dir_compilation" ] ;
#        then mkdir -p $dir_compilation
#        fi

abigen  --abi ./../abi/metaMaster/metaMaster.abi \
        --bin ./../abi/metaMaster/metaMaster.bin \
        --type metaMaster \
        --pkg metamaster \
        --out ./../abi/metaMaster/metaMaster.go

abigen  --abi ./../abi/saleAuction/saleAuction.abi \
        --bin ./../abi/saleAuction/saleAuction.bin \
        --type saleAuction \
        --pkg saleauction \
        --out ./../abi/saleAuction/saleAuction.go
abigen  --abi ./../abi/salePlain/salePlain.abi \
        --bin ./../abi/salePlain/salePlain.bin \
        --type salePlain \
        --pkg saleplain \
        --out ./../abi/salePlain/salePlain.go

