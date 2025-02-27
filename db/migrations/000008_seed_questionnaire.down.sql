BEGIN;

-- Remove exclusions first since they depend on answers
DELETE FROM exclusions WHERE answer_ulid IN (
    '01JKZMRP780464QR5GRKKJFE38',
    '01JKZMRSR7698EXDTAPWDQ0ETM',
    '01JKZMRTB10CH599A7EYGSH9XW',
    '01JKZMRTMFSB4WTEXT75PJ6ZWZ',
    '01JKZMRTXWGKFN49AB53PJDPBP',
    '01JKZMRV78D9Q3XBMFKBAP4A36',
    '01JKZMRVT37ERXGE0WFYARWM4Y',
    '01JKZMRW3GSKKW7WBZ49SF5F53',
    '01JKZMRWCX2F37QSSVF5PWA66P',
    '01JKZMRWPBX62ZYX5QDZKR6D9Q'
);

-- Remove answer flows since they depend on answers
DELETE FROM answers_flow WHERE answer_ulid IN (
    '01JKZMRNXVQ4QSJ548FBSV1GJX',
    '01JKZMRP780464QR5GRKKJFE38',
    '01JKZMRPGPWW4GF1B55BBS3R9Z',
    '01JKZMRPT3JY4TE22R00EQA8JN',
    '01JKZMRQ3GS5731G5MJ9Y79DMH',
    '01JKZMRQCYDJCSNCBRFBSGHPHG',
    '01JKZMRQPBA9KNNAV8GZ9B5WPV',
    '01JKZMRQZRXD1QXQD7FSXDDYNW',
    '01JKZMRR9582G1T895FHET9F9Q',
    '01JKZMRRJK98K7BW5W1SR0JNCF',
    '01JKZMRRW0Q28YT4XNR16GP89N',
    '01JKZMRS5DFMC97TRGKKWETT0E',
    '01JKZMRSETM7J6C8J5JKC3BCH1',
    '01JKZMRSR7698EXDTAPWDQ0ETM',
    '01JKZMRT1MEB1360JZMF55ZF83',
    '01JKZMRTB10CH599A7EYGSH9XW',
    '01JKZMRTMFSB4WTEXT75PJ6ZWZ',
    '01JKZMRTXWGKFN49AB53PJDPBP',
    '01JKZMRV78D9Q3XBMFKBAP4A36',
    '01JKZMRVGNECK1R25XJJ33PCDD',
    '01JKZMRVT37ERXGE0WFYARWM4Y',
    '01JKZMRW3GSKKW7WBZ49SF5F53',
    '01JKZMRWCX2F37QSSVF5PWA66P',
    '01JKZMRWPBX62ZYX5QDZKR6D9Q',
    '01JKZMRWZRHMSSZ5WQ2FDPX8A6'
);

-- Remove answer recommendations
DELETE FROM answer_recommendations WHERE answer_ulid IN (
    '01JKZMRQCYDJCSNCBRFBSGHPHG',
    '01JKZMRQPBA9KNNAV8GZ9B5WPV',
    '01JKZMRQZRXD1QXQD7FSXDDYNW',
    '01JKZMRR9582G1T895FHET9F9Q',
    '01JKZMRRJK98K7BW5W1SR0JNCF',
    '01JKZMRRW0Q28YT4XNR16GP89N',
    '01JKZMRS5DFMC97TRGKKWETT0E',
    '01JKZMRSETM7J6C8J5JKC3BCH1'
);

-- Remove recommendations
DELETE FROM recommendations WHERE ulid IN (
    '01JKZMRYET3167075ZNHT9GVVV',
    '01JKZMRYR8BAZ54SAWE8Y54ZZP',
    '01JKZMRZ1N1JCKH23D2KTZ5PD6',
    '01JKZMRZB23CFG018169XNETC2'
);

-- Remove products
DELETE FROM products WHERE ulid IN (
    '01JKZMRX95JNSGHY74NK3YYHBP',
    '01JKZMRXJJZ302H3YRSHK635TH',
    '01JKZMRXW09S50H04M7ZEZ93TG',
    '01JKZMRY5DY4MYNZRRNN2982FS'
);

-- Remove answers
DELETE FROM answers WHERE ulid IN (
    '01JKZMRNXVQ4QSJ548FBSV1GJX',
    '01JKZMRP780464QR5GRKKJFE38',
    '01JKZMRPGPWW4GF1B55BBS3R9Z',
    '01JKZMRPT3JY4TE22R00EQA8JN',
    '01JKZMRQ3GS5731G5MJ9Y79DMH',
    '01JKZMRQCYDJCSNCBRFBSGHPHG',
    '01JKZMRQPBA9KNNAV8GZ9B5WPV',
    '01JKZMRQZRXD1QXQD7FSXDDYNW',
    '01JKZMRR9582G1T895FHET9F9Q',
    '01JKZMRRJK98K7BW5W1SR0JNCF',
    '01JKZMRRW0Q28YT4XNR16GP89N',
    '01JKZMRS5DFMC97TRGKKWETT0E',
    '01JKZMRSETM7J6C8J5JKC3BCH1',
    '01JKZMRSR7698EXDTAPWDQ0ETM',
    '01JKZMRT1MEB1360JZMF55ZF83',
    '01JKZMRTB10CH599A7EYGSH9XW',
    '01JKZMRTMFSB4WTEXT75PJ6ZWZ',
    '01JKZMRTXWGKFN49AB53PJDPBP',
    '01JKZMRV78D9Q3XBMFKBAP4A36',
    '01JKZMRVGNECK1R25XJJ33PCDD',
    '01JKZMRVT37ERXGE0WFYARWM4Y',
    '01JKZMRW3GSKKW7WBZ49SF5F53',
    '01JKZMRWCX2F37QSSVF5PWA66P',
    '01JKZMRWPBX62ZYX5QDZKR6D9Q',
    '01JKZMRWZRHMSSZ5WQ2FDPX8A6'
);

-- Remove questions
DELETE FROM questions WHERE ulid IN (
    '01JKZMRKJHW7MF4HKN0DAY0PAR',
    '01JKZMRKVYH2SQXS5R9EXQQYWR',
    '01JKZMRM5B09QNCFK46B5YRFP0',
    '01JKZMRMES1XQ5X3NKDTF5B0NZ',
    '01JKZMRMR63Z3P8RAW74GYWRS4',
    '01JKZMRN1KJSFJGP1693RX7PZ1',
    '01JKZMRNB0GK5V1K2CNXFJTAAT',
    '01JKZMRNMEE3V9HT8RHBV25QJR'
);

COMMIT;
