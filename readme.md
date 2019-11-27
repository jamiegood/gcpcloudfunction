## Summary
Learning GCP cloud functions see https://cloud.google.com/functions/docs/tutorials/pubsub

## Deploying cloud function
gcloud functions deploy HelloPubSub --runtime go111 --trigger-topic ToBigQuery

## Send message to function
gcloud pubsub topics publish ToBigQuery --message YOUR_NAME2

Insert into bq
bq query --use_legacy_sql=false "INSERT hwtechradar.mytable (author, title) VALUES ('HG WellsXX', 'Time MachineXX')" 


Publish message to pubsub
gcloud pubsub topics publish JamieTestTopic --message '{ "author": "wednesday author", "title": "wednesday title" }'  