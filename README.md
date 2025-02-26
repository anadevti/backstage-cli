# Backstage CLI

Um CLI desenvolvido em Go para gerenciar o desenvolvimento e a implantação do Backstage, incluindo integração com o ArgoCD.

---

## Funcionalidades

- **Subir o Backstage localmente**: Executa `yarn dev` para iniciar o Backstage no ambiente de desenvolvimento.
- **Sincronizar com o ArgoCD**: Sincroniza a aplicação Backstage com o ArgoCD para implantações em Kubernetes. (A desenvolver)

---

## Pré-requisitos

- **Go**: Instale o Go (versão 1.20 ou superior). Siga as instruções em [golang.org](https://golang.org/doc/install).
- **Yarn**: Instale o Yarn para gerenciar dependências do Backstage. Siga as instruções em [yarnpkg.com](https://yarnpkg.com/getting-started/install).
- **ArgoCD CLI** (a desenvolver): Se você quiser usar a integração com o ArgoCD, instale o CLI do ArgoCD. Siga as instruções em [ArgoCD Documentation](https://argo-cd.readthedocs.io/en/stable/cli_installation/).

---

## Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/backstage-cli.git
   cd backstage-cli
   ```

2. Compile o projeto:
   ```bash
   go build -o backstage-cli
   ```

3. Mova o binário para um diretório no seu PATH (opcional):
   ```bash
   sudo mv backstage-cli /usr/local/bin/
   ```

---

## Como Usar

### Subir o Backstage localmente

Para subir o Backstage localmente usando `yarn dev`, execute:

```bash
./backstage-cli up --path /caminho/para/o/projeto/clonado_backstage
```

- **`--path`**: Caminho para o diretório do projeto Backstage. Se omitido, o CLI usará o diretório atual (`.`).

### Sincronizar com o ArgoCD (A DESENVOLVER)

Para sincronizar a aplicação Backstage com o ArgoCD, execute:

```bash
./backstage-cli sync
```

- Certifique-se de que o `argocd` CLI está instalado e configurado.
- O comando sincroniza a aplicação chamada `backstage` no ArgoCD.

---

## Exemplo de Uso

1. Navegue até o diretório do seu projeto clonado do Backstage:
   ```bash
   cd /caminho/para/seu/projeto/backstage_clone
   ```

2. Suba o Backstage localmente:
   ```bash
   backstage-cli up
   ```

3. Após fazer alterações no Backstage, sincronize com o ArgoCD:
   ```bash
   backstage-cli sync
   ```

---

## Estrutura do Projeto

```
backstage-cli/
├── main.go          # Código principal do CLI
├── go.mod          # Dependências do Go
├── go.sum          # Checksum das dependências
└── README.md       # Este arquivo
```

---

## Contribuição

Contribuições são muito bem-vindas! Siga os passos abaixo:

1. Faça um fork do projeto.
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`).
3. Commit suas mudanças (`git commit -m 'Adicionando nova feature'`).
4. Push para a branch (`git push origin feature/nova-feature`).
5. Abra um Pull Request.

---

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

## Contato

Se você tiver dúvidas ou sugestões, sinta-se à vontade para entrar em contato:

- **Nome**: Ana Carolyne
- **Email**: anacarolayne777@gmail.com