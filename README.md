# Hidrate-se !["golang version"](https://img.shields.io/badge/golang-v1.15.6-green) !["wails version"](https://img.shields.io/badge/wails-v1.10.1-blue) !["wails version"](https://img.shields.io/badge/status-developing-red) [![Go Report Card](https://goreportcard.com/badge/github.com/joaovicdsantos/hidrate-se)](https://goreportcard.com/report/github.com/joaovicdsantos/hidrate-se)

Aplicativo feito para aqueles que tem uma dificuldade em beber corretamente água ao longo do dia. O Hidrate-se notifica o usuário em um certo período de tempo para que ele possa atingir sua meta diária (que é devidamente calculada pelo software). 

## Aspectos Técnicos

Foi desenvolvido em golang usando o framework [__wails__](https://github.com/wailsapp/wails). No visual, optei por utilizar somente HTML, CSS e JS com um webpack gerenciando tudo, por se tratar de um aplicativo simples. Para o gerenciamento das notifcações foi usado o pacote [__beeep__](https://github.com/gen2brain/beeep). 

## Tarefas:
- [ ] Fazer o programa rodar em segundo plano
- [ ] Implementar um ícone funcional na tray
- [ ] Adicionar parâmetro para o tempo de funcionamento do programa
- [ ] Persistir os dados passados até o usuário desejar alterá-los
- [x] Cria uma logo e indentidade visual

## Screenshots

#### Tela inicial:
!["tela inical"](https://i.ibb.co/jk1cRWF/hidrate-se1.png)
#### Após começar:
!["notificacao"](https://i.ibb.co/TBVTSrC/hidrate-se2.png)
#### Notificação:
!["iniciado"](https://i.ibb.co/XCF9d7d/hidrate-se3.png)
