If you've opened this, let's start with a "Thank you".
Questo è il mio primo progetto, che spero sia l'inizio di una lunga serie.
Ho lavorato mesi nell'incertezza più totale per arrivare a questo punto e sono successe molte cose in questa finestra di tempo, ma torniamo al progetto.
---------------------------------------------------------------------------------------------------------------------------------------------------------
Questo è KEEPER, un'architettura a microservizi pensata per le PMI. Nel nostro caso, la PMI sarà la DGAuto, un concessionario multimarca presente nel 
territorio salentino. Al momento è composto da 4 filiali, due per provincia tra Lecce e Brindisi. Ogni filiale ha un minimo di:
-1 Assistente;
-2 Venditori;
-1 Responsabile.
-1 Responsabile territoriale per provincia;
-Ovviamente il personale della C-Board sarà unico e avrà accesso completo al 90% delle funzioni.

Il tutto è organizzato in una struttura piramidale per gestire le autorizzazioni e la sicurezza del sistema, in modo che personale meno qualificato o
con meno autorizzazioni non possa effettuare pesanti modifiche al sistema.

--------------------------------------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------SCELTA TECNOLOGIE---------------------------------------------------------------
--------------------------------------------------------------------------------------------------------------------------------------------------------
Per il Back-End si è deciso di utilizzare le seguenti tecnologie:
-GoLang: sia per portfolio sia per efficenza ed efficacia di questo linguaggio per questi applicativi;
-GORM: il più diffuso ORM per GoLang, preferito a "sqlc" per poter mettere mano prima su un ORM puro rispetto a un query-generator;
-Python3: verrà utilizzato per il microservizio secondario relativo alle notifiche;
-PeeWee: ORM per Python, preferito a SQLAlchemy per semplicità, data l'interazione limitata che ha Python con il resto del sistema. Nonostante ciò esso
         non limita la scalabilità del progetto;
-PostgreSQL: DataBase relazionale scelto per il salvataggio dei vari dati. Perfetto per grandi volumi di dati e query complesse ma soprattutto con carichi
             di lavoro misti in scrittura/lettura, il vero e proprio Framework che possiede lo fa preferire per scalabilità a rivali come MySQL;
-RestAPI: sarà lo standard utilizzato per la comunicazione tra servizi, con un formato dati JSON;

