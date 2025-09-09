-- Criar banco de dados e usar
CREATE DATABASE IF NOT EXISTS empresa;
USE empresa;

-- ==========================
-- TABELAS
-- ==========================

-- 1. Colaborador
CREATE TABLE colaborador (
    id_colaborador INT NOT NULL AUTO_INCREMENT,
    nome VARCHAR(100) NOT NULL,
    cpf VARCHAR(14) UNIQUE DEFAULT NULL,
    cargo VARCHAR(50) DEFAULT NULL,
    setor VARCHAR(50) DEFAULT NULL,
    senha VARCHAR(60) NOT NULL, -- suporta bcrypt/argon2
    status ENUM('Ativo','Ausente','Férias','Home Office') DEFAULT 'Ativo',
    email VARCHAR(100) UNIQUE DEFAULT NULL,
    ramal VARCHAR(20) DEFAULT NULL,
    habilidades TEXT,
    PRIMARY KEY (id_colaborador)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 2. Projeto
CREATE TABLE projeto (
    id_projeto INT NOT NULL AUTO_INCREMENT,
    nome VARCHAR(100) NOT NULL,
    tipo ENUM('Software','Security') NOT NULL,
    status ENUM('Em Andamento','Atrasado','No Prazo') DEFAULT 'Em Andamento',
    progresso INT DEFAULT '0',
    descricao TEXT,
    PRIMARY KEY (id_projeto)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 3. Equipe do Projeto
CREATE TABLE equipe_projeto (
    id_projeto INT NOT NULL,
    id_colaborador INT NOT NULL,
    funcao VARCHAR(50) DEFAULT NULL,
    PRIMARY KEY (id_projeto,id_colaborador),
    KEY (id_colaborador),
    CONSTRAINT fk_eqp_projeto FOREIGN KEY (id_projeto) REFERENCES projeto (id_projeto),
    CONSTRAINT fk_eqp_colaborador FOREIGN KEY (id_colaborador) REFERENCES colaborador (id_colaborador)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 4. Tarefa
CREATE TABLE tarefa (
    id_tarefa INT NOT NULL AUTO_INCREMENT,
    titulo VARCHAR(100) NOT NULL,
    id_projeto INT DEFAULT NULL,
    prazo DATE DEFAULT NULL,
    prioridade ENUM('Baixa','Média','Alta','Urgente') DEFAULT 'Média',
    status ENUM('To-Do','Doing','Done','Urgente','Pendente','Planejado','Em Andamento') DEFAULT 'To-Do',
    responsavel INT DEFAULT NULL,
    PRIMARY KEY (id_tarefa),
    KEY (id_projeto),
    KEY (responsavel),
    CONSTRAINT fk_tarefa_projeto FOREIGN KEY (id_projeto) REFERENCES projeto (id_projeto),
    CONSTRAINT fk_tarefa_responsavel FOREIGN KEY (responsavel) REFERENCES colaborador (id_colaborador)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 5. Backup
CREATE TABLE backup (
    id_backup INT NOT NULL AUTO_INCREMENT,
    data TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    mantido_ate DATE DEFAULT NULL,
    responsavel INT DEFAULT NULL,
    PRIMARY KEY (id_backup),
    KEY (responsavel),
    CONSTRAINT fk_backup_colaborador FOREIGN KEY (responsavel) REFERENCES colaborador (id_colaborador)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 6. Chamado
CREATE TABLE chamado (
    id_chamado INT NOT NULL AUTO_INCREMENT,
    setor VARCHAR(50) DEFAULT NULL,
    prioridade ENUM('Baixa','Média','Alta','Urgente') DEFAULT 'Média',
    status ENUM('Aberto','Em Andamento','Resolvido','Fechado') DEFAULT 'Aberto',
    tecnico_designado INT DEFAULT NULL,
    PRIMARY KEY (id_chamado),
    KEY (tecnico_designado),
    CONSTRAINT fk_chamado_tecnico FOREIGN KEY (tecnico_designado) REFERENCES colaborador (id_colaborador)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 7. Registro de Ponto
CREATE TABLE registroponto (
    id_ponto INT NOT NULL AUTO_INCREMENT,
    id_colaborador INT DEFAULT NULL,
    entrada DATETIME DEFAULT NULL,
    saida DATETIME DEFAULT NULL,
    PRIMARY KEY (id_ponto),
    KEY (id_colaborador),
    CONSTRAINT fk_ponto_colaborador FOREIGN KEY (id_colaborador) REFERENCES colaborador (id_colaborador)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ==========================
-- INSERTS EXEMPLO
-- ==========================

-- Colaboradores com hash bcrypt fictício (use password_hash no app real)

INSERT INTO colaborador (nome, cpf, cargo, setor, senha, status, email, ramal, habilidades) VALUES
('José Silva', '111.111.111-11', 'Desenvolvedor', 'TI', '$2y$10$abcdefghijklmnopqrstuvABCDEFGHIJKLMNOPQRSTUV123456', 'Ativo', 'jose@empresa.com', '101', 'PHP, MySQL'),
('Maria Souza', '222.222.222-22', 'Analista de Segurança', 'Segurança', '$2y$10$mnopqrABCDEFGHIJKLMNopqrstuvWXYZabcdefghijkl12345', 'Ativo', 'maria@empresa.com', '102', 'Red Team, Pentest'),
('Ana Lima', '333.333.333-33', 'Suporte', 'Suporte', '$2y$10$qrstuvABCDEFGHIJKLMNOPQRSTUVabcdefghijklmnop123456', 'Ativo', 'ana@empresa.com', '103', 'Atendimento, Helpdesk'),
('Carlos Pereira', '444.444.444-44', 'Gerente de Projetos', 'PMO', '$2y$10$uvwxyzABCDEFGHIJKLMNopqrstuvWXYZabcdefghijkl78901', 'Ativo', 'carlos@empresa.com', '104', 'Gestão, Scrum'),
('Paulo Almeida', '555.555.555-55', 'Desenvolvedor', 'TI', '$2y$10$abcdefGHIJKLMNOPQRSTUVwxyz1234567890abcdefghijklmn', 'Ativo', 'paulo@empresa.com', '105', 'Java, Spring'),
('tiago','322.233.112-50','gestor','gestor','$2b$10$BfHJ84Z8E4/ZGwfFcsWSx.GzTw4jtdhV58n4WH2xKgUfB.GrA52du','Ativo','tiagodasilva@gmail.com','2131','nada');
-- Projetos
INSERT INTO projeto (nome, tipo, status, progresso, descricao) VALUES
('Sistema Interno', 'Software', 'Em Andamento', 20, 'Desenvolvimento do sistema interno da empresa'),
('Teste de Segurança', 'Security', 'No Prazo', 50, 'Auditoria e pentest do sistema legado');

-- Equipes
INSERT INTO equipe_projeto (id_projeto, id_colaborador, funcao) VALUES
(1, 1, 'Desenvolvedor'),
(1, 4, 'Gerente'),
(1, 5, 'Desenvolvedor'),
(2, 2, 'Analista de Segurança'),
(2, 4, 'Gerente');

-- Tarefas
INSERT INTO tarefa (titulo, id_projeto, prazo, prioridade, status, responsavel) VALUES
('Criar base de dados', 1, '2025-09-10', 'Alta', 'To-Do', 1),
('Implementar login', 1, '2025-09-12', 'Média', 'To-Do', 5),
('Teste de vulnerabilidade', 2, '2025-09-15', 'Urgente', 'To-Do', 2);

-- Chamados
INSERT INTO chamado (setor, prioridade, status, tecnico_designado) VALUES
('Infraestrutura', 'Alta', 'Aberto', 3),
('Desenvolvimento', 'Média', 'Em Andamento', 1);

-- Backups
INSERT INTO backup (mantido_ate, responsavel) VALUES
('2025-09-10', 3),
('2025-09-12', 1);

-- Registro de Ponto
INSERT INTO registroponto (id_colaborador, entrada, saida) VALUES
(1, '2025-09-01 08:00:00', '2025-09-01 17:00:00'),
(2, '2025-09-01 09:00:00', '2025-09-01 18:00:00'),
(3, '2025-09-01 08:30:00', '2025-09-01 17:30:00');
