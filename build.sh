cd cmd/assistant || exit
go build -o gpt-assistant

cd ../../cmd/chat || exit
go build -o gpt-chat
