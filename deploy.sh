#!/bin/bash

gcloud functions deploy go-http-function \
--env-vars-file=.env.yaml.local \
--gen2 \
--runtime=go119 \
--region=asia-northeast1 \
--source=. \
--entry-point=GetOGP \
--trigger-http \
--allow-unauthenticated

exit 0