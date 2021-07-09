create table if not exists account
(
    id           serial primary key,
    email        text,
    username     text,
    firstName    text,
    lastName     text,
    profileImage text,
    password     text,
    type         text
);
insert into account(email, username, firstName, lastName, profileImage, password, type)
values ('1', 'amyadams20', 'Amy', 'Adams', 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/a41cf22c-578d-456f-b0c4-90d5c017eab1', 1, ''),
       ('2', 'billybilly', 'Billy', 'Childs', 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/781069ca-2f84-46b9-8597-abaa8344f574', 1, ''),
       ('3', 'sammy', 'Sam', 'Smith', 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/c7e5a0cb-19b6-44f0-a28a-bba219271344', 1, '');

create table if not exists destination
(
    id          serial primary key,
    name        text,
    country     text,
    description text not null default '',
    imageName   text,
    imageData   text,
    latitude    float8,
    longitude   float8
);
insert into destination(name, country, description, imageName, imageData, latitude, longitude)
values ('Paris', 'France',
        'Paris (nicknamed the \"City of light\") is the capital city of France, and the largest city in France. The area is 105 square kilometres (41 square miles), and around 2.15 million people live there. If suburbs are counted, the population of the Paris area rises to 12 million people.',
        'https://res.cloudinary.com/vietname/image/upload/v1625311603/paris.jpg', '', 48.855014, 2.341231),
       ('Tokyo', 'Japan',
        'Tokyo (東京, Tōkyō) is Japan''s capital and the world''s most populous metropolis. It is also one of Japan''s 47 prefectures, consisting of 23 central city wards and multiple cities, towns and villages west of the city center. The Izu and Ogasawara Islands are also part of Tokyo.',
        'https://res.cloudinary.com/vietname/image/upload/v1625311650/japan.jpg', '', 35.67988, 139.7695),
       ('New York', 'US',
        'New York City has been described as the cultural, financial, and media capital of the world, significantly influencing commerce, entertainment, research, technology, education, politics, tourism, art, fashion, and sports.',
        'https://res.cloudinary.com/vietname/image/upload/v1625433115/newyork.jpg', '', 40.71592, -74.0055);

create table if not exists destination_detail
(
    id             serial primary key,
    destination_id int,
    photoUrl       text
);
alter table destination_detail
    add constraint fk_destination_id foreign key (destination_id) references destination (id);
insert into destination_detail(destination_id, photoUrl)
values (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/7156c3c6-945e-4284-a796-915afdc158b5'),
       (1, 'https://letsbuildthatapp-videos.s3-us-west-2.amazonaws.com/b1642068-5624-41cf-83f1-3f6dff8c1702'),
       (1, 'https://letsbuildthatapp-videos.s3-us-west-2.amazonaws.com/6982cc9d-3104-4a54-98d7-45ee5d117531'),
       (1, 'https://letsbuildthatapp-videos.s3-us-west-2.amazonaws.com/2240d474-2237-4cd3-9919-562cd1bb439e'),
       (2, 'https://letsbuildthatapp-videos.s3-us-west-2.amazonaws.com/bda78f98-0050-4f40-9a2c-0160d9043c06'),
       (2, 'https://letsbuildthatapp-videos.s3-us-west-2.amazonaws.com/ef73f53a-bf37-406f-8175-64fdb102e5b8'),
       (2, 'https://letsbuildthatapp-videos.s3-us-west-2.amazonaws.com/88e646e8-df59-46ee-82ba-bb7f86015373'),
       (3, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/be17bf36-b7ad-43e0-8689-5f4c3a20e9e0'),
       (3, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/f2bed6fe-5fcc-4899-9e5a-7be1c7526074');

create table if not exists destination_rank
(
    id             serial primary key,
    destination_id int    not null default 0,
    vote           float8 not null default 5,
    total_of_vote  int    not null default 1
);
alter table destination_rank
    add constraint fk_destination_id foreign key (destination_id) references destination (id);
insert into destination_rank(destination_id, vote)
values (1, 5),
       (2, 5),
       (3, 5);

create table if not exists restaurant
(
    id          serial primary key,
    name        text,
    imageName   text,
    city        text,
    country     text,
    description text,
    created_by  int
);
alter table restaurant
    add constraint fk_created_by foreign key (created_by) references
        account (id);
insert into restaurant (name, imageName, city, country, description, created_by)
values ('Japan''s Finest Tapas', 'https://res.cloudinary.com/vietname/image/upload/v1625370568/newyork.jpg', 'Tokyo', 'Japan', 'The traditional cuisine of Japan, washoku (和食), lit. \"Japanese eating\" (or kappō (ja:割烹)), is based on rice with miso soup and other dishes; there is an emphasis on seasonal ingredients. Side dishes often consist of fish, pickled vegetables, and vegetables cooked in broth.', 1),
       ('Bar & Grill', 'https://res.cloudinary.com/vietname/image/upload/v1625370700/newyork.jpg', 'New York', 'US', 'Oaklands Park''s favourite with great food and beverages, kids area & much more. Make a reservation now!', 2);

create table if not exists restaurant_rank
(
    id            serial primary key,
    restaurant_id int    not null default 0,
    vote          float8 not null default 5,
    total_of_vote int    not null default 1
);
alter table restaurant_rank
    add constraint fk_restaurant_id foreign key (restaurant_id)
        references restaurant (id);
insert into restaurant_rank (restaurant_id, vote, total_of_vote)
values (1, 5, 1),
       (2, 5, 1);

create table if not exists restaurant_dishes
(
    id            serial primary key,
    restaurant_id int not null,
    name          text,
    price         text,
    numPhotos     int,
    photo         text
);
alter table restaurant_dishes
    add constraint fk_restaurant_id foreign key
        (restaurant_id) references restaurant (id);
insert into restaurant_dishes(restaurant_id, name, price, numPhotos, photo)
values (1, 'Dragon Rolls', '$12.99', 9,
        'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/e2f3f5d4-5993-4536-9d8d-b505d7986a5c'),
       (1, 'Sashimi', '$19.99', 15,
        'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/a4d85eff-4c79-4141-a0d6-761cca48eae1'),
       (1, 'Combination Rolls', '$22.99', 5,
        'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/20a6783b-3de7-4e58-9e22-bcc6a43b6df6'),
       (1, 'Bento Boxes', '$8.49', 1,
        'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/0d1d2e79-2f10-4cfd-82da-a1c2ab3638d2'),
       (1, 'Fresh Rolls', '$24.59', 3,
        'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/3a352f87-3dc1-4fa7-affe-fb12fa8691fe');

create table if not exists restaurant_photos
(
    id            serial primary key,
    restaurant_id int not null,
    photoUrl      text,
    description   text,
    created_on    text
);
alter table restaurant_photos
    add constraint fk_restaurant_id foreign key
        (restaurant_id) references restaurant (id);
insert into restaurant_photos(restaurant_id, photoUrl, description, created_on)
values (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/e2f3f5d4-5993-4536-9d8d-b505d7986a5c',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid). In some bars and restaurants in Spain and across the globe, tapas have evolved into a more sophisticated cuisine.',
        '30/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/a4d85eff-4c79-4141-a0d6-761cca48eae1',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/20a6783b-3de7-4e58-9e22-bcc6a43b6df6',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/0d1d2e79-2f10-4cfd-82da-a1c2ab3638d2',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/3923d237-3931-44e5-836f-5de40ec04b31',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/254c0418-2b55-4a2b-b530-a31a9799c7d5',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/fa20d064-b6d7-4df9-8f44-0f25f6ee5a19',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/a441d22b-5324-4444-8ddf-22b99128838c',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/6b5d013b-dc3b-4e5e-93d9-ec932f42aead',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/a6de1d65-8fa3-4674-a6ce-a207b8f86b15',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/5c6bc68c-a8a1-42ac-ab3a-947927826807',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/a5e83c0c-c815-4129-bfd4-17e73fa1da78',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/f6ee5fb7-b21b-42c1-b1d8-a455742d0247',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/c22e8d9e-10f2-4559-8c81-375491295e84',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/3a352f87-3dc1-4fa7-affe-fb12fa8691fe',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/8ca76521-1f52-4043-8b86-d2a573342daf',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021'),
       (1, 'https://letsbuildthatapp-videos.s3.us-west-2.amazonaws.com/73f69749-f986-46ac-9b8b-d7b1d42bddc5',
        'A tapa (Spanish pronunciation: [ˈtapa]) is an appetizer or snack in Spanish cuisine. Tapas may be cold (such as mixed olives and cheese) or hot (such as chopitos, which are battered, fried baby squid)',
        '28/06/2021');

create table if not exists restaurant_reviews
(
    id            serial primary key,
    restaurant_id int not null,
    account_id    int not null,
    rating        int,
    text          text
);
alter table restaurant_reviews
    add constraint fk_restaurant_id foreign key
        (restaurant_id) references restaurant (id);
alter table restaurant_reviews
    add constraint fk_account_id foreign key
        (account_id) references account (id);

insert into restaurant_reviews(restaurant_id, account_id, rating, text)
values (1, 1, 4, 'They are always here on time, never leave early and adhere to all company break times.'),
       (1, 2, 5,
        'You regularly follow up with customers to ensure they are having a great experience with the product, and as a result, you have brought in many new orders.'),
       (1, 3, 2,
        'You always gather all of the information and facts to make a decision, which benefits the entire team.'),
       (1, 1, 3, 'Many customers have complained that you have been rude with them.');

--
--
-- drop table restaurant_reviews, restaurant_photos,
--     restaurant_dishes, restaurant_rank, restaurant,
--     destination_rank, destination_detail, destination, account;