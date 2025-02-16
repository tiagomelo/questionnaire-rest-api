BEGIN;

INSERT INTO questions (ulid, label, text) VALUES 
('01JKZMRKJHW7MF4HKN0DAY0PAR', '1', 'Do you have difficulty getting or maintaining an erection?'),
('01JKZMRKVYH2SQXS5R9EXQQYWR', '2', 'Have you tried any of the following treatments before?'),
('01JKZMRM5B09QNCFK46B5YRFP0', '2a', 'Was the Viagra or Sildenafil product you tried before effective?'),
('01JKZMRMES1XQ5X3NKDTF5B0NZ', '2b', 'Was the Cialis or Tadalafil product you tried before effective?'),
('01JKZMRMR63Z3P8RAW74GYWRS4', '2c', 'Which is your preferred treatment?'),
('01JKZMRN1KJSFJGP1693RX7PZ1', '3', 'Do you have, or have you ever had, any heart or neurological conditions?'),
('01JKZMRNB0GK5V1K2CNXFJTAAT', '4', 'Do any of the listed medical conditions apply to you?'),
('01JKZMRNMEE3V9HT8RHBV25QJR', '5', 'Are you taking any of the following drugs?');

INSERT INTO answers (ulid, question_ulid, text, next_question_ulid, previous_question_ulid) VALUES 
-- Q1
('01JKZMRNXVQ4QSJ548FBSV1GJX', '01JKZMRKJHW7MF4HKN0DAY0PAR', 'Yes', '01JKZMRKVYH2SQXS5R9EXQQYWR', NULL),
('01JKZMRP780464QR5GRKKJFE38', '01JKZMRKJHW7MF4HKN0DAY0PAR', 'No', NULL, NULL),

-- Q2
('01JKZMRPGPWW4GF1B55BBS3R9Z', '01JKZMRKVYH2SQXS5R9EXQQYWR', 'Viagra or Sildenafil', '01JKZMRM5B09QNCFK46B5YRFP0', '01JKZMRKJHW7MF4HKN0DAY0PAR'),
('01JKZMRPT3JY4TE22R00EQA8JN', '01JKZMRKVYH2SQXS5R9EXQQYWR', 'Cialis or Tadalafil', '01JKZMRMES1XQ5X3NKDTF5B0NZ', '01JKZMRKJHW7MF4HKN0DAY0PAR'),
('01JKZMRQ3GS5731G5MJ9Y79DMH', '01JKZMRKVYH2SQXS5R9EXQQYWR', 'Both', '01JKZMRMR63Z3P8RAW74GYWRS4', '01JKZMRKJHW7MF4HKN0DAY0PAR'),
('01JKZMRQCYDJCSNCBRFBSGHPHG', '01JKZMRKVYH2SQXS5R9EXQQYWR', 'None of the above', '01JKZMRN1KJSFJGP1693RX7PZ1', '01JKZMRKJHW7MF4HKN0DAY0PAR'),

-- Q2a
('01JKZMRQPBA9KNNAV8GZ9B5WPV', '01JKZMRM5B09QNCFK46B5YRFP0', 'Yes', '01JKZMRN1KJSFJGP1693RX7PZ1', '01JKZMRKVYH2SQXS5R9EXQQYWR'),
('01JKZMRQZRXD1QXQD7FSXDDYNW', '01JKZMRM5B09QNCFK46B5YRFP0', 'No', '01JKZMRN1KJSFJGP1693RX7PZ1', '01JKZMRKVYH2SQXS5R9EXQQYWR'),

-- Q2b
('01JKZMRR9582G1T895FHET9F9Q', '01JKZMRMES1XQ5X3NKDTF5B0NZ', 'Yes', '01JKZMRN1KJSFJGP1693RX7PZ1', '01JKZMRKVYH2SQXS5R9EXQQYWR'),
('01JKZMRRJK98K7BW5W1SR0JNCF', '01JKZMRMES1XQ5X3NKDTF5B0NZ', 'No', '01JKZMRN1KJSFJGP1693RX7PZ1', '01JKZMRKVYH2SQXS5R9EXQQYWR'),

-- Q2c
('01JKZMRRW0Q28YT4XNR16GP89N', '01JKZMRMR63Z3P8RAW74GYWRS4', 'Viagra or Sildenafil', '01JKZMRN1KJSFJGP1693RX7PZ1', '01JKZMRKVYH2SQXS5R9EXQQYWR'),
('01JKZMRS5DFMC97TRGKKWETT0E', '01JKZMRMR63Z3P8RAW74GYWRS4', 'Cialis or Tadalafil', '01JKZMRN1KJSFJGP1693RX7PZ1', '01JKZMRKVYH2SQXS5R9EXQQYWR'),
('01JKZMRSETM7J6C8J5JKC3BCH1', '01JKZMRMR63Z3P8RAW74GYWRS4', 'None of the above', '01JKZMRN1KJSFJGP1693RX7PZ1', '01JKZMRKVYH2SQXS5R9EXQQYWR'),

-- Q3
('01JKZMRSR7698EXDTAPWDQ0ETM', '01JKZMRN1KJSFJGP1693RX7PZ1', 'Yes', NULL, '01JKZMRKVYH2SQXS5R9EXQQYWR'),
('01JKZMRT1MEB1360JZMF55ZF83', '01JKZMRN1KJSFJGP1693RX7PZ1', 'No', '01JKZMRNB0GK5V1K2CNXFJTAAT', '01JKZMRKVYH2SQXS5R9EXQQYWR'),

-- Q4
('01JKZMRTB10CH599A7EYGSH9XW', '01JKZMRNB0GK5V1K2CNXFJTAAT', 'Significant liver problems (such as cirrhosis of the liver) or kidney problems', NULL, '01JKZMRN1KJSFJGP1693RX7PZ1'),
('01JKZMRTMFSB4WTEXT75PJ6ZWZ', '01JKZMRNB0GK5V1K2CNXFJTAAT', 'Currently prescribed GTN, Isosorbide mononitrate, Isosorbide dinitrate , Nicorandil (nitrates) or Rectogesic ointment', NULL, '01JKZMRN1KJSFJGP1693RX7PZ1'),
('01JKZMRTXWGKFN49AB53PJDPBP', '01JKZMRNB0GK5V1K2CNXFJTAAT', 'Abnormal blood pressure (lower than 90/50 mmHg or higher than 160/90 mmHg)', NULL, '01JKZMRN1KJSFJGP1693RX7PZ1'),
('01JKZMRV78D9Q3XBMFKBAP4A36', '01JKZMRNB0GK5V1K2CNXFJTAAT', 'Condition affecting your penis (such as Peyronie''s Disease, previous injuries or an inability to retract your foreskin)', NULL, '01JKZMRN1KJSFJGP1693RX7PZ1'),
('01JKZMRVGNECK1R25XJJ33PCDD', '01JKZMRNB0GK5V1K2CNXFJTAAT', 'I don''t have any of these conditions', '01JKZMRNMEE3V9HT8RHBV25QJR', '01JKZMRN1KJSFJGP1693RX7PZ1'),

-- Q5
('01JKZMRVT37ERXGE0WFYARWM4Y', '01JKZMRNMEE3V9HT8RHBV25QJR', 'Alpha-blocker medication such as Alfuzosin, Doxazosin, Tamsulosin, Prazosin, Terazosin or over-the-counter Flomax', NULL, '01JKZMRNB0GK5V1K2CNXFJTAAT'),
('01JKZMRW3GSKKW7WBZ49SF5F53', '01JKZMRNMEE3V9HT8RHBV25QJR', 'Riociguat or other guanylate cyclase stimulators (for lung problems)', NULL, '01JKZMRNB0GK5V1K2CNXFJTAAT'),
('01JKZMRWCX2F37QSSVF5PWA66P', '01JKZMRNMEE3V9HT8RHBV25QJR', 'Saquinavir, Ritonavir or Indinavir (for HIV)', NULL, '01JKZMRNB0GK5V1K2CNXFJTAAT'),
('01JKZMRWPBX62ZYX5QDZKR6D9Q', '01JKZMRNMEE3V9HT8RHBV25QJR', 'Cimetidine (for heartburn)', NULL, '01JKZMRNB0GK5V1K2CNXFJTAAT'),
('01JKZMRWZRHMSSZ5WQ2FDPX8A6', '01JKZMRNMEE3V9HT8RHBV25QJR', 'I don''t take any of these drugs', NULL, '01JKZMRNB0GK5V1K2CNXFJTAAT');

INSERT INTO products (ulid, name, identifier) VALUES 
('01JKZMRX95JNSGHY74NK3YYHBP', 'Sildenafil 50mg', 'sildenafil_50'),
('01JKZMRXJJZ302H3YRSHK635TH', 'Sildenafil 100mg', 'sildenafil_100'),
('01JKZMRXW09S50H04M7ZEZ93TG', 'Tadalafil 10mg', 'tadalafil_10'),
('01JKZMRY5DY4MYNZRRNN2982FS', 'Tadalafil 20mg', 'tadalafil_20');

INSERT INTO recommendations (ulid, product_ulid) VALUES 
-- Sildenafil 50mg
('01JKZMRYET3167075ZNHT9GVVV', '01JKZMRX95JNSGHY74NK3YYHBP'),
-- Sildenafil 100mg
('01JKZMRYR8BAZ54SAWE8Y54ZZP', '01JKZMRXJJZ302H3YRSHK635TH'),
-- Tadalafil 10mg
('01JKZMRZ1N1JCKH23D2KTZ5PD6', '01JKZMRXW09S50H04M7ZEZ93TG'),
-- Tadalafil 20mg
('01JKZMRZB23CFG018169XNETC2', '01JKZMRY5DY4MYNZRRNN2982FS');

INSERT INTO answer_recommendations (answer_ulid, recommendation_ulid) VALUES 
-- Q2 'None of the above', Sildenafil 50mg
('01JKZMRQCYDJCSNCBRFBSGHPHG', '01JKZMRYET3167075ZNHT9GVVV'),
-- Q2 'None of the above', Tadalafil 10mg
('01JKZMRQCYDJCSNCBRFBSGHPHG', '01JKZMRZ1N1JCKH23D2KTZ5PD6'),
-- Q2a 'Yes', Sildenafil 50mg
('01JKZMRQPBA9KNNAV8GZ9B5WPV', '01JKZMRYET3167075ZNHT9GVVV'),
-- Q2a 'No', Tadalafil 20mg
('01JKZMRQZRXD1QXQD7FSXDDYNW', '01JKZMRZB23CFG018169XNETC2'),
-- Q2b 'Yes', Tadalafil 10mg
('01JKZMRR9582G1T895FHET9F9Q', '01JKZMRZ1N1JCKH23D2KTZ5PD6'),
-- Q2b 'No', Sildenafil 100mg
('01JKZMRRJK98K7BW5W1SR0JNCF', '01JKZMRYR8BAZ54SAWE8Y54ZZP'),
-- Q2c, 'Viagra or Sildenafil', Sildenafil 100mg
('01JKZMRRW0Q28YT4XNR16GP89N', '01JKZMRYR8BAZ54SAWE8Y54ZZP'),
-- Q2c 'Cialis or Tadalafil', Tadalafil 20mg
('01JKZMRS5DFMC97TRGKKWETT0E', '01JKZMRZB23CFG018169XNETC2'),
-- Q2c 'None of the above', Sildenafil 100mg
('01JKZMRSETM7J6C8J5JKC3BCH1', '01JKZMRYR8BAZ54SAWE8Y54ZZP'),
-- Q2c 'None of the above', Tadalafil 20mg
('01JKZMRSETM7J6C8J5JKC3BCH1', '01JKZMRZB23CFG018169XNETC2');

INSERT INTO answers_flow (answer_ulid, previous_answer_ulid, next_question_ulid) VALUES
-- Q1 'Yes', no previous answer, next question is Q2
('01JKZMRNXVQ4QSJ548FBSV1GJX', NULL, '01JKZMRKVYH2SQXS5R9EXQQYWR'),
-- Q1 'No', no previous answer neither next questions, exit
('01JKZMRP780464QR5GRKKJFE38', NULL, NULL),

-- Q2 'Viagra or Sildenafil', previous answer is Q1 'Yes', next question is Q2a
('01JKZMRPGPWW4GF1B55BBS3R9Z', '01JKZMRNXVQ4QSJ548FBSV1GJX', '01JKZMRM5B09QNCFK46B5YRFP0'),
-- Q2 'Cialis or Tadalafil', previous answer is Q1 'Yes', next question is Q2b
('01JKZMRPT3JY4TE22R00EQA8JN', '01JKZMRNXVQ4QSJ548FBSV1GJX', '01JKZMRMES1XQ5X3NKDTF5B0NZ'),
-- Q2 'Both', previous answer is Q1 'Yes', next question is Q2c
('01JKZMRQ3GS5731G5MJ9Y79DMH', '01JKZMRNXVQ4QSJ548FBSV1GJX', '01JKZMRMR63Z3P8RAW74GYWRS4'),
-- Q2 'None of the above', previous answer is Q1 'Yes', next question is Q3
('01JKZMRQCYDJCSNCBRFBSGHPHG', '01JKZMRNXVQ4QSJ548FBSV1GJX', '01JKZMRN1KJSFJGP1693RX7PZ1'),

-- Q2a 'Yes', previous answer is Q2 'Viagra or Sildenafil', next question is Q3
('01JKZMRQPBA9KNNAV8GZ9B5WPV', '01JKZMRPGPWW4GF1B55BBS3R9Z', '01JKZMRN1KJSFJGP1693RX7PZ1'),
-- Q2a 'No', previous answer is Q2 'Viagra or Sildenafil', next question is Q3
('01JKZMRQZRXD1QXQD7FSXDDYNW', '01JKZMRPGPWW4GF1B55BBS3R9Z', '01JKZMRN1KJSFJGP1693RX7PZ1'),

-- Q2b 'Yes', previous answer is Q2 'Cialis or Tadalafil', next question is Q3
('01JKZMRR9582G1T895FHET9F9Q', '01JKZMRPT3JY4TE22R00EQA8JN', '01JKZMRN1KJSFJGP1693RX7PZ1'),
-- Q2b 'No', previous answer is Q2 'Cialis or Tadalafil', next question is Q3
('01JKZMRRJK98K7BW5W1SR0JNCF', '01JKZMRPT3JY4TE22R00EQA8JN', '01JKZMRN1KJSFJGP1693RX7PZ1'),

-- Q2c 'Viagra or Sildenafil', previous answer is Q2 'Both', next question is Q3
('01JKZMRRW0Q28YT4XNR16GP89N', '01JKZMRQ3GS5731G5MJ9Y79DMH', '01JKZMRN1KJSFJGP1693RX7PZ1'),
-- Q2c 'Cialis or Tadalafil',  previous answer is Q2 'Both', next question is Q3
('01JKZMRS5DFMC97TRGKKWETT0E', '01JKZMRQ3GS5731G5MJ9Y79DMH', '01JKZMRN1KJSFJGP1693RX7PZ1'),
-- Q2c 'None of the above', previous answer is Q2 'Both', next question is Q3
('01JKZMRSETM7J6C8J5JKC3BCH1', '01JKZMRQ3GS5731G5MJ9Y79DMH', '01JKZMRN1KJSFJGP1693RX7PZ1'),

-- Q3 'Yes', previous answer is Q2 'None of the above', no next question, exit
('01JKZMRSR7698EXDTAPWDQ0ETM', '01JKZMRQCYDJCSNCBRFBSGHPHG', NULL),

-- Q3 'Yes', previous answer is Q2a 'Yes', no next question, exit
('01JKZMRSR7698EXDTAPWDQ0ETM', '01JKZMRQPBA9KNNAV8GZ9B5WPV', NULL),
-- Q3 'Yes', previous answer is Q2a 'No', no next question, exit
('01JKZMRSR7698EXDTAPWDQ0ETM', '01JKZMRQZRXD1QXQD7FSXDDYNW', NULL),

-- Q3 'Yes', previous answer is Q2b 'Yes', no next question, exit
('01JKZMRSR7698EXDTAPWDQ0ETM', '01JKZMRR9582G1T895FHET9F9Q', NULL),
-- Q3 'Yes', previous answer is Q2b 'No', no next question, exit
('01JKZMRSR7698EXDTAPWDQ0ETM', '01JKZMRRJK98K7BW5W1SR0JNCF', NULL),

-- Q3 'Yes', previous answer is Q2c 'Viagra or Sildenafil', no next question, exit
('01JKZMRSR7698EXDTAPWDQ0ETM', '01JKZMRRW0Q28YT4XNR16GP89N', NULL),
-- Q3 'Yes', previous answer is Q2c 'Cialis or Tadalafil', no next question, exit
('01JKZMRSR7698EXDTAPWDQ0ETM', '01JKZMRS5DFMC97TRGKKWETT0E', NULL),
-- Q3 'Yes', previous answer is Q2c 'None of the above', no next question, exit
('01JKZMRSR7698EXDTAPWDQ0ETM', '01JKZMRSETM7J6C8J5JKC3BCH1', NULL),

-- Q3 'No', previous answer is Q2 'None of the above', next question is Q4
('01JKZMRT1MEB1360JZMF55ZF83', '01JKZMRQCYDJCSNCBRFBSGHPHG', '01JKZMRNB0GK5V1K2CNXFJTAAT'),
-- Q3 'No', previous answer is Q2a 'Yes', next question is Q4
('01JKZMRT1MEB1360JZMF55ZF83', '01JKZMRQPBA9KNNAV8GZ9B5WPV', '01JKZMRNB0GK5V1K2CNXFJTAAT'),
-- Q3 'No', previous answer is Q2a 'No', next question is Q4
('01JKZMRT1MEB1360JZMF55ZF83', '01JKZMRQZRXD1QXQD7FSXDDYNW', '01JKZMRNB0GK5V1K2CNXFJTAAT'),
-- Q3 'No', previous answer is Q2b 'Yes', next question is Q4
('01JKZMRT1MEB1360JZMF55ZF83', '01JKZMRR9582G1T895FHET9F9Q', '01JKZMRNB0GK5V1K2CNXFJTAAT'),
-- Q3 'No', previous answer is Q2b 'No', next question is Q4
('01JKZMRT1MEB1360JZMF55ZF83', '01JKZMRRJK98K7BW5W1SR0JNCF', '01JKZMRNB0GK5V1K2CNXFJTAAT'),
-- Q3 'No', previous answer is Q2c 'Viagra or Sildenafil', next question is Q4
('01JKZMRT1MEB1360JZMF55ZF83', '01JKZMRRW0Q28YT4XNR16GP89N', '01JKZMRNB0GK5V1K2CNXFJTAAT'),
-- Q3 'No', previous answer is Q2c 'Cialis or Tadalafil', next question is Q4
('01JKZMRT1MEB1360JZMF55ZF83', '01JKZMRS5DFMC97TRGKKWETT0E', '01JKZMRNB0GK5V1K2CNXFJTAAT'),
-- Q3 'No', previous answer is Q2c 'None of the above', next question is Q4
('01JKZMRT1MEB1360JZMF55ZF83', '01JKZMRSETM7J6C8J5JKC3BCH1', '01JKZMRNB0GK5V1K2CNXFJTAAT'),

-- Q4 'Significant liver problems (such as cirrhosis of the liver) or kidney problems', previous answer is Q3 'No', no next question, exit
('01JKZMRTB10CH599A7EYGSH9XW', '01JKZMRT1MEB1360JZMF55ZF83', NULL),
-- Q4 'Currently prescribed GTN, Isosorbide mononitrate, Isosorbide dinitrate , Nicorandil (nitrates) or Rectogesic ointment', previous answer is Q3 'No', no next question, exit
('01JKZMRTMFSB4WTEXT75PJ6ZWZ', '01JKZMRT1MEB1360JZMF55ZF83', NULL),
-- Q4 'Abnormal blood pressure (lower than 90/50 mmHg or higher than 160/90 mmHg)', previous answer is Q3 'No', no next question, exit
('01JKZMRTXWGKFN49AB53PJDPBP', '01JKZMRT1MEB1360JZMF55ZF83', NULL),
-- Q4 'Condition affecting your penis (such as Peyronie''s Disease, previous injuries or an inability to retract your foreskin', previous answer is Q3 'No', no next question, exit
('01JKZMRV78D9Q3XBMFKBAP4A36', '01JKZMRT1MEB1360JZMF55ZF83', NULL),
-- Q4 'I don''t have any of these conditions', previous answer is Q3 'No', next question is Q5
('01JKZMRVGNECK1R25XJJ33PCDD', '01JKZMRT1MEB1360JZMF55ZF83', '01JKZMRNMEE3V9HT8RHBV25QJR'),

-- Q5 'Alpha-blocker medication such as Alfuzosin, Doxazosin, Tamsulosin, Prazosin, Terazosin or over-the-counter Flomax', previous answer is ''I don''t have any of these conditions'', no next question, exit
('01JKZMRVT37ERXGE0WFYARWM4Y', '01JKZMRVGNECK1R25XJJ33PCDD', NULL),
-- Q5 'Riociguat or other guanylate cyclase stimulators (for lung problems)', previous answer is ''I don''t have any of these conditions'', no next question, exit
('01JKZMRW3GSKKW7WBZ49SF5F53', '01JKZMRVGNECK1R25XJJ33PCDD', NULL),
-- Q5 'Saquinavir, Ritonavir or Indinavir (for HIV)', previous answer is ''I don''t have any of these conditions'', no next question, exit
('01JKZMRWCX2F37QSSVF5PWA66P', '01JKZMRVGNECK1R25XJJ33PCDD', NULL),
-- Q5 'Cimetidine (for heartburn)', previous answer is ''I don''t have any of these conditions'', no next question, exit
('01JKZMRWPBX62ZYX5QDZKR6D9Q', '01JKZMRVGNECK1R25XJJ33PCDD', NULL),
-- Q5 'I don''t take any of these drugs', previous answer is ''I don''t have any of these conditions'', no next question, exit
('01JKZMRWZRHMSSZ5WQ2FDPX8A6', '01JKZMRVGNECK1R25XJJ33PCDD', NULL);

INSERT INTO exclusions (answer_ulid, reason) VALUES
-- Q1 'No', no products available
('01JKZMRP780464QR5GRKKJFE38', 'No products available'),
-- Q3 'Yes', no products available
('01JKZMRSR7698EXDTAPWDQ0ETM', 'No products available'),
-- Q4 'Significant liver problems (such as cirrhosis of the liver) or kidney problems', no products available
('01JKZMRTB10CH599A7EYGSH9XW', 'No products available'),
-- Q4 'Currently prescribed GTN, Isosorbide mononitrate, Isosorbide dinitrate , Nicorandil (nitrates) or Rectogesic ointment', no products available
('01JKZMRTMFSB4WTEXT75PJ6ZWZ', 'No products available'),
-- Q4 'Abnormal blood pressure (lower than 90/50 mmHg or higher than 160/90 mmHg)', no products available
('01JKZMRTXWGKFN49AB53PJDPBP', 'No products available'),
-- Q4 'Condition affecting...', no products available
('01JKZMRV78D9Q3XBMFKBAP4A36', 'No products available'),
-- Q5 'Alpha-blocker medication...', no products available
('01JKZMRVT37ERXGE0WFYARWM4Y', 'No products available'),
-- Q5 'Riociguat or other guanylate...', no products available
('01JKZMRW3GSKKW7WBZ49SF5F53', 'No products available'),
-- Q5 'Saquinavir, Ritonavir or Indinavir...', no products available
('01JKZMRWCX2F37QSSVF5PWA66P', 'No products available'),
-- Q5 'Cimetidine (for heartburn)', no products available
('01JKZMRWPBX62ZYX5QDZKR6D9Q', 'No products available');

COMMIT;