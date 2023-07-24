CREATE TABLE ad_space (
    ad_space_id INT AUTO_INCREMENT PRIMARY KEY,
    ad_space_name VARCHAR(255) NOT NULL,
    ad_space_price  DECIMAL(10,2) NOT NULL, 
    is_ad_available BOOLEAN,
    auction_time VARCHAR(255) NOT NULL
);


CREATE TABLE bidders (
    bidder_id INT NOT NULL AUTO_INCREMENT,
    bidder_name varchar(255) NOT NULL,
    bidder_budget int(10),
    bid_time VARCHAR(255) NOT NULL,
    PRIMARY KEY (bidder_id)
);

CREATE TABLE bids (
    id INT AUTO_INCREMENT PRIMARY KEY,
    bidder_id INT NOT NULL,
    ad_space_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    bid_time DATETIME NOT NULL,
    FOREIGN KEY (bidder_id) REFERENCES bidders (bidder_id),
    FOREIGN KEY (ad_space_id) REFERENCES ad_space (ad_space_id)
);

INSERT INTO ad_space (ad_space_name, ad_space_price, is_ad_available, auction_time)
VALUES ('Ad Space 1', 100.00, true, '2023-07-25 15:00:00'),
       ('Ad Space 2', 75.50, true, '2023-07-26 12:30:00'),
       ('Ad Space 3', 50.25, false, '2023-07-27 18:45:00');

INSERT INTO bidders (bidder_name, bidder_budget, bid_time)
VALUES ('Bidder 1', 500.00, '2023-07-25 15:00:00'),
       ('Bidder 2', 1000.00, '2023-07-27 18:45:00'),
       ('Bidder 3', 750.00, '2023-07-27 18:45:00');

INSERT INTO bids (bidder_id, ad_space_id, amount, bid_time)
VALUES (1, 1, 80.00, '2023-07-25 14:30:00'),
       (2, 2, 70.00, '2023-07-26 12:35:00'),
       (3, 1, 90.50, '2023-07-25 15:15:00');

