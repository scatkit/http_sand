--create table raydium_markets (
--  id integer primary key,
--  base_mint TEXT NOT NULL,
--  quote_mint TEXT NOT NULL
--);
--
--insert into raydium_markets(
--  id,
--  base_mint,
--  quote_mint 
--) VALUES (
--  1,
--  '5m9AhhwMnosop8CAWdSJDNfLLkg3EFkX5sJ3zAuMpump',
--  'So11111111111111111111111111111111111111112'
--);

--create table pairs(
--  id integer primary key,
--  name text not NULL,
--  raydium_market_id integer not null,
--  foreign key (raydium_market_id)
-- references raydium_markets(id)
--);
 
--insert into pairs(
--  id,
--  name,
--  raydium_market_id
--) VALUES(
--  1,
--  'WILDOUT:SOL',
--  1
--);

select raydium_market_id from pairs where name = 'WILDOUT:SOL';
