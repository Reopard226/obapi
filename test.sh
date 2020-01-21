set -e
touch coverage.txt
echo mode: "atomic" > coverage.txt
export ENVKEY=V64jBasy94hdsqV1F95s-uFADhGBvhXaKQvm9
for d in $(go list ./... | grep -v -e internal/mock -e cmd/api/server); do
    go test -race -coverprofile=profile.out -covermode=atomic "$d"
    if [ -f profile.out ]; then
        echo "$(tail -n +2 profile.out)" > profile.out
        cat profile.out >> coverage.txt
        #rm profile.out
    fi
done

sed -i '/^$/d' coverage.txt