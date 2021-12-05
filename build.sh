#! /bin/bash

BASE=$PWD
BUILD=$BASE/build
REACT=$BASE/botonawire-react
API=$BASE/botonawire-backend

if [ ! -d $BUILD ] 
then
    mkdir $BUILD
else
    rm -rf $BUILD/*
fi

cd $REACT
npm run build
mv ./build $BUILD/static

cd $API
export CGO_ENABLED=0
go build
cp botonawire $BUILD

cd $BASE
docker build -t botonawire-server:latest .
