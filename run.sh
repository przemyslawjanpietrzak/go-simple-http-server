SERVER_NAME="server.go"
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "${DIR}"
cwd=$(pwd)
SOURCE="${BASH_SOURCE[0]}"
echo "${SOURCE}"
eval "go run ${DIR}/${SERVER_NAME} ${cwd} $1"
