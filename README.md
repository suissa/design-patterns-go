# Design Patterns Go

Coleção de Design Patterns em Go.

## 1. Abstract Factory
Exemplo: Uma fábrica de móveis que produz conjuntos completos (cadeira, mesa, sofá) em diferentes estilos (moderno, clássico).

Como funciona: Você escolhe o estilo ("moderno"), e a fábrica entrega todos os móveis compatíveis entre si, sem precisar saber detalhes de construção.

## 2. Builder
Exemplo: Montar um computador personalizado.

Como funciona: Você escolhe peças passo a passo (processador, RAM, placa de vídeo) e o "montador" (builder) constrói a máquina final com base nessas escolhas.

## 3. Prototype
Exemplo: Cópia de documentos em uma impressora.

Como funciona: Você faz uma cópia idêntica de um documento original sem precisar recriá-lo do zero.

## 4. Singleton
Exemplo: Um logger global em um aplicativo.

Como funciona: Garante que todas as partes do sistema usem a mesma instância para registrar eventos, evitando múltiplas instâncias conflitantes.

## 5. Adapter
Exemplo: Adaptador de tomada para viajar entre países.

Como funciona: Converte a interface de um plugue (padrão brasileiro) para funcionar em uma tomada europeia.

## 6. Bridge
Exemplo: Controles remotos para TVs de marcas diferentes.

Como funciona: O controle remoto (abstração) funciona com qualquer TV (implementação), independente da marca (Samsung, LG).

## 7. Composite
Exemplo: Sistema de arquivos em um computador.

Como funciona: Pastas (composite) e arquivos (leaf) são tratados da mesma forma: você pode mover, copiar ou excluir ambos.

## 8. Decorator
Exemplo: Personalização de um café em uma cafeteria.

Como funciona: Você adiciona "decoradores" (leite, chantilly, chocolate) a um café básico, incrementando seu custo e sabor.

## 9. Facade
Exemplo: Interface de um home theater.

Como funciona: Um único botão "Assistir Filme" liga a TV, o som e ajusta a iluminação, escondendo a complexidade dos subsistemas.

## 10. Flyweight
Exemplo: Renderização de caracteres em um editor de texto.

Como funciona: Caracteres repetidos (como a letra "A") compartilham a mesma representação em memória, economizando recursos.

## 11. Proxy
Exemplo: Serviço de streaming (Netflix).

Como funciona: O proxy carrega apenas a miniatura do vídeo inicialmente, e o conteúdo completo é baixado apenas quando você clica para assistir.

## 12. Chain of Responsibility
Exemplo: Sistema de suporte técnico.

Como funciona: Uma solicitação passa por níveis de atendimento (bot de chat → atendente → supervisor) até ser resolvida.

## 13. Command
Exemplo: Controles de um drone por controle remoto.

Como funciona: Cada botão (decolar, pousar) encapsula uma ação como um objeto, permitindo desfazer/refazer operações.

## 14. Iterator
Exemplo: Navegação em uma playlist do Spotify.

Como funciona: Você usa "próxima música" ou "voltar" para percorrer a lista sem precisar saber como ela é armazenada.

## 15. Mediator
Exemplo: Sistema de chat em um aplicativo.

Como funciona: O mediador gerencia a comunicação entre usuários, evitando que eles se comuniquem diretamente.

## 16. Memento
Exemplo: Salvamento automático em jogos.

Como funciona: O jogo salva o estado atual (nível, vida, itens) para que você possa restaurá-lo se falhar.

## 17. Observer
Exemplo: Notificações de redes sociais.

Como funciona: Quando alguém curte sua foto, todos os inscritos (observadores) recebem uma notificação automática.

## 18. Visitor
Exemplo: Auditoria em uma loja de departamentos.

Como funciona: Um auditor (visitor) verifica estoque, preços e validade de produtos sem alterar as classes dos objetos analisados.