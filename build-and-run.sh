dir=$(pwd)
if [ ! -f go-bindata ]
then
	go build github.com/jteeuwen/go-bindata/go-bindata
	if [[ $? != 0 ]]
	then
		exit
	fi
fi
./go-bindata -o src/perfmongo_app/resources.go src/page \
&& go build perfmongo_app \
&& ./perfmongo_app
