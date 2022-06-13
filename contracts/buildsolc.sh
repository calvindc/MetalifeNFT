#!/bin/sh

# the version of solc is Version: 0.8.1+commit.df193b15.Linux.g++
# solc-select install 0.8.1
# solc-select use 0.8.1

dir_bin_metaMaster=./abi/metaMaster
dir_bin_saleAuction=./abi/saleAuction
dir_bin_salePlain=./abi/salePlain
if [ ! -d "$dir_bin_metaMaster" ] ;
        then mkdir -p $dir_bin_metaMaster
        fi
if [ ! -d "$dir_bin_salePlain" ] ;
        then mkdir -p $dir_bin_salePlain
        fi
if [ ! -d "$dir_bin_saleAuction" ] ;
        then mkdir -p $dir_bin_saleAuction
        fi


solc --base-path ./ --allow-paths ./node_moudules/@openzeppelin --optimize --optimize-runs 200 --overwrite --abi --bin -o ./abi/metaMaster  ./metaMaster.sol

solc --base-path ./ --allow-paths ./node_moudules/@openzeppelin --optimize --optimize-runs 200 --overwrite --abi --bin -o ./abi/salePlain  ./salePlain.sol

solc --base-path ./ --allow-paths ./node_moudules/@openzeppelin --optimize --optimize-runs 200 --overwrite --abi --bin -o ./abi/saleAuction  ./saleAuction.sol
