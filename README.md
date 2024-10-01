<p align="center">
  <img src="https://github.com/user-attachments/assets/e3a36db8-7ed3-491d-89cf-a0d04ed0d415" alt="Descrição da Imagem">
</p>


# Go Essentials - LinuxTips
Neste Repo vou colocar minha experiência com o curso Go Essentials da Linuxtips serão três dias de conteúdo e m projeto final que será submetido ao fim do curso, a idéia é documentar brevemente sobre os aprendizados de cada dia.

## Day 1
Rodamos nosso primeiro programa Go, tivemos contato com uma biblioteca da linguagem, a poderosa "fmt" além de aprendemos sobre tipos e formatacao de print
Achei bem interessante a forma como os packages funcionam, para fixar criei um pequeno gráfico com a ideia basica da coisa.
Inclusive algumas idéias começaram a vir sobre formas de práticar mais durante o curso, talvez criar um pacote com uma entidade Logger ou então estudar e criar teste unitários para os dias de estudo que passaremos por aqui.

```plaintext
+-------------------------+
|         mathutils       |
+-------------------------+
| + Add(a int, b int) int | <------------------+
| + Multiply(a int, b int)|                    |
+-------------------------+                    |
                                               |
                                               |
                                               |
            +----------------------------------+
            |   main.go (exemplo)              |
            +----------------------------------+
            |                                  |
            | - import "mathutils"             |
            |                                  |
            | - func main() {                  |
            |     result := mathutils.Add(1, 2)|
            |     fmt.Println(result)          |
            | }                                |
            +----------------------------------+
```

## Day 2
No segundo dia do curso Go Essentials, mergulhamos no desenvolvimento de uma poderosa calculadora e aprofundamos nosso entendimento sobre funções de controle em Go. Aprendemos a utilizar condicionais, como if/else, e loops, especialmente o for, que são fundamentais para controlar o fluxo de execução dos programas. Também introduzimos o comando switch, que se mostrou uma alternativa clara e concisa para lidar com múltiplas condições.

No exercício de criar a calculadora, fui um pouco além e desenvolvi uma estrutura de package para organizar as funções.

Estou aproveitando a integração nativa da linguagem para aprofundar meus conhecimentos sobre testes. Decidi me desafiar a implementar testes sempre que possível em diferentes partes do código, pois acredito que os testes são essenciais para um desenvolvimento seguro e confiável.