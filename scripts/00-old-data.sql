insert into places ("name",address,link,slug,inverted_logo) values
	 ('Дом печати','Екатеринбург, ул. Ленина, 49','https://tele-club.ru/dompechati','dompechati',true),
	 ('U-bar','Екатеринбург, ул. Добролюбова, 3В','https://vk.com/u_bar_music_loft','ubar',false),
	 ('TeleClub','Екатеринбург, ул. Карьерная, 16','https://tele-club.ru','teleclub',true),
	 (' ','Екатеринбург, пр. Ленина, 11а','','street',false),
	 ('Свобода Концерт Холл','Екатеринбург, ул. Черкасская, 12','https://svoboda-ekb.ru/','svoboda',true),
	 ('Корчма "Пристанище"','Екатеринбург, ул. Бебеля, 124','https://vk.link/tavern_ekb','korchma',false),
	 ('Syndrome Bar','Екатеринбург, ул. 8 Марта, д.13, вход со двора','https://syndromebar.ru/','syndrome',false),
	 ('Бар Ц','Екатеринбург, пер. Центральный рынок, 6','https://tele-club.ru/c','barc',false),
	 ('Клуб "7|44"','Екатеринбург, пр. Ленина, 48','https://vk.com/musquartal','744',true);

insert into gigs (dt,tm,place,"desc",url) values
	 ('2019-02-03','19:00:00',(select id from places where slug = 'dompechati'),'Полуфинал конкурса Emergenza 2019',''),
	 ('2019-02-09','20:00:00',(select id from places where slug = 'ubar'),'',''),
	 ('2019-04-20','14:30:00',(select id from places where slug = 'teleclub'),'Финал конкурса Emergenza 2019',''),
	 ('2019-07-13','23:50:00',(select id from places where slug = 'ubar'),'',''),
	 ('2019-10-18','23:40:00',(select id from places where slug = 'ubar'),'',''),
	 ('2022-05-28','19:30:00',(select id from places where slug = 'street'),'Выступление в рамках фестиваля Библионочь. Концерт «Народ, вера и рок-н-ролл»',''),
	 ('2022-06-24','20:00:00',(select id from places where slug = 'svoboda'),'Выступим на верхней сцене в рамках фестиваля Ural Music Night. Начало программы в 21:00. Наш выход в 1:30',''),
	 ('2022-07-30','21:30:00',(select id from places where slug = 'korchma'),'',''),
	 ('2022-10-08','20:00:00',(select id from places where slug = 'syndrome'),'Большой сольный концерт',''),
	 ('2022-10-23','18:00:00',(select id from places where slug = 'barc'),'Принимаем участие в новом сезоне Поколение-РОК 2022-2023!',''),
	 ('2022-11-27','18:00:00',(select id from places where slug = 'barc'),'Еще один раунд фестиваля Поколение-РОК',''),
	 ('2022-12-16','18:00:00',(select id from places where slug = '744'),'Участвуем в отборочном туре на фестиваль Yletaй',''),
	 ('2022-12-23','20:00:00',(select id from places where slug = '744'),'Участвуем в качестве приглашенных гостей на сольном концерте пост-рок группы WAY',''),
	 ('2023-05-21','17:30:00',(select id from places where slug = 'svoboda'),'Поборемся за право выступить на фестивале "Музыка Свободы"','https://vk.com/otbormuzikasvobody'),
	 ('2023-06-02','20:00:00',(select id from places where slug = 'svoboda'),'Будем частью большого музыкального праздника. Приходи затусить в компании отличных ребят.','https://vk.com/dvizhpartysvoboda');
