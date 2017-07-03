#!/bin/sh
SOURCE_DIR=$GOPATH/src/MadridMas/server
if [ ! -d "$GOPATH" ]; then
  echo "Error: \$GOPATH is not a directory"
  exit 1
fi
if [ ! -d "$SOURCE_DIR" ]; then
  echo "Error: $SOURCE_DIR does not exist"
  exit 1
fi
if [ ! -d "$SOURCE_DIR/.git" ]; then
  echo "Error: $SOURCE_DIR is not a git repository"
  exit 1
fi
if [ $# -eq 0 ]; then
  echo "Nothing to build. Specify at least one of 'oracle'."
  exit 0
fi

for EXECUTABLE in "$@"; do
  cd $SOURCE_DIR/$EXECUTABLE
  go build
  if [ $? -ne 0 ]; then
    echo "Error: build of '$EXECUTABLE' failed"
    exit 1
  else
    echo "Build of '$EXECUTABLE' succeeded"
  fi
done

