protoc -Inextdate/proto --go_opt=module=github.com/Vova4o/todogrpc --go_out=. --go-grpc_opt=module=github.com/Vova4o/todogrpc --go-grpc_out=. nextdate/proto/*.proto

Собрать gRPC ручками или через Makefile

Частичная копия проекта Todo на HTTP, выполненная на gRPC

Сервер выполнен про принципу "чистой архитектуры", что позволяет вносить изменения в части кода не прибегая к изменениям в других частях кода.
