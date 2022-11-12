# Sniffer de pacotes para N nodos

Este é um trabalho acadêmico desenvolvido para a disciplina de Redes de computadores I, com o intuito de explorar e comparar os protocolos de transporte UDP, TCP e SCTP.

A aplicação consiste na comunicação de N nodos que irão compartilhar entre si, um sniffer da sua rede local, por exemplo o endereço A sniffará sua rede local e enviará o resultado para os demais nodos.

Cada nodo terá um servidor rodando localmente, pois ele receberá as requisições dos outro endereços, sniffará a rede e enviará os dados sniffados. Da mesma maneira, cada nodo criará um cliente para comunicar-se com os outros endereços, solicitando os pacotes.

É importante lembrar que todas máquinas precisam estar rodando me paralelo para que a comunicação entre eleas flue corretamente, pois se uma das máquinas não estiver executando, não será possível obter os pacotes da mesma.

## Diagramas
### Diagrama de funcionamento das comunicações entre os nodos
![Untitled-2022-10-23-2122](https://user-images.githubusercontent.com/47281681/201498244-0113ca51-e663-4197-8e10-d1338afc0b95.png)

### Diagrama de funcionamento interno do nodo
![Untitled-2022-10-23-2122](https://user-images.githubusercontent.com/47281681/201498445-19a03027-229c-4956-a685-03f95cd0210b.png)

# Adicionar ou remover nodos
O programa executa de acordo com os nodos configurados no arquivo `ADICIONAR O NOME DEPOIS.json`, onde devemos ter o endereço de ip da máquina e a porta que será utilizada para comunicação. É importante ressaltar que, por não ter um servidor centralizado, a maquina que estiver executando a aplicação, necessita ter o endereço registrado nesse arquivo.

# Como executar

Acessando a pasta raiz do repositório execute o seguinte comando, informando o protocolo desejado:
**OBS:** Somente os protocolos UDP, TCP e SCTP foram implementados.

````
go run . -protocol <NOME_DO_PROTOCOLO>
````

## Análise do comportamento para os protocolos
