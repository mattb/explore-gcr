helm upgrade excited-gorilla chart --set image.tag=$(git rev-list --max-count 1 HEAD)
