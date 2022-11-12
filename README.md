## Sniffer de pacotes para N nodos

Este é um trabalho acadêmico desenvolvido para a disciplina de Redes de computadores I, com o intuito de explorar e comparar os protocolos de transporte UDP, TCP e SCTP.

A aplicação consiste na comunicação de N nodos que irão compartilhar entre si, um sniffer da sua rede local, por exemplo o endereço A sniffará sua rede local e enviará o resultado para os demais nodos.


## Como executar

Acessando a pasta raiz do repositório execute o seguinte comando, informando o protocolo desejado:
**OBS:** Somente os protocolos UDP, TCP e SCTP foram implementados.

````
go run . -protocol {NOME_DO_PROTOCOLO}
