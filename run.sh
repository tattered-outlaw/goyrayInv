#!/bin/bash
docker run -v $(pwd)/images:/usr/src/app/images -it --rm --name goray-running goray