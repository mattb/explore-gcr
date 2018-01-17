helm upgrade excited-gorilla chart --set image.tag=$(gcloud container images list-tags gcr.io/mattb-k8s/github-mattb-explore-gcr --format json | jq -r .[0].tags[0])
