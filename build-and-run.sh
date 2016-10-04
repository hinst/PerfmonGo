dir=$(pwd)
go build github.com/jteeuwen/go-bindata/go-bindata \
&& ./go-bindata -o src/perfmongo_app/resources.go src/page \
&& go build perfmongo_app \
&& ./perfmongo_app
