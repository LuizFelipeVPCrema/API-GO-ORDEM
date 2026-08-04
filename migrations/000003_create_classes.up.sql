CREATE TABLE classes (
    id INTEGER PRIMARY KEY,

    codigo TEXT NOT NULL UNIQUE
        CHECK (
            codigo IN (
                'COMBATENTE',
                'ESPECIALISTA',
                'OCULTISTA'
            )
        ),

    nome TEXT NOT NULL UNIQUE,

    ordem_exibicao INTEGER NOT NULL UNIQUE,

    ativa INTEGER NOT NULL DEFAULT 1
        CHECK (ativa IN (0, 1)),

    criado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    atualizado_em DATETIME NOT NULL
        DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO classes (
    id,
    codigo,
    nome,
    ordem_exibicao
)
VALUES
    (
        1,
        'COMBATENTE',
        'Combatente',
        1
    ),
    (
        2,
        'ESPECIALISTA',
        'Especialista',
        2
    ),
    (
        3,
        'OCULTISTA',
        'Ocultista',
        3
    );