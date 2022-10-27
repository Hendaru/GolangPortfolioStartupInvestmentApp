-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Oct 23, 2022 at 03:40 AM
-- Server version: 5.7.34
-- PHP Version: 7.4.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bwstartup`
--

-- --------------------------------------------------------

--
-- Table structure for table `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int(50) UNSIGNED NOT NULL,
  `user_id` int(50) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `short_description` varchar(100) DEFAULT NULL,
  `description` text,
  `goal_amount` int(50) DEFAULT NULL,
  `current_amount` int(50) DEFAULT NULL,
  `backer_count` int(50) DEFAULT NULL,
  `perks` text,
  `slug` varchar(50) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `campaigns`
--

INSERT INTO `campaigns` (`id`, `user_id`, `name`, `short_description`, `description`, `goal_amount`, `current_amount`, `backer_count`, `perks`, `slug`, `created_at`, `updated_at`) VALUES
(9, 13, 'campaign 1 update', 'diskripsi campaign 1 update', 'penjelasan yg panjang lebar update', 10, 0, 0, 'keuntungan satu, kemudian yang dua, dan ketiga update', 'campaign-1-13', '2022-10-15 10:47:09', '2022-10-15 10:48:49'),
(10, 14, 'campaign 2', 'diskripsi campaign 2', 'penjelasan yg panjang lebar', 10000, 0, 0, 'keuntungan satu, kemudian yang dua, dan ketiga', 'campaign-2-14', '2022-10-15 14:34:52', '2022-10-15 14:34:52'),
(11, 14, 'campaign 2', 'diskripsi campaign 2', 'penjelasan yg panjang lebar', 10000, 0, 0, 'keuntungan satu, kemudian yang dua, dan ketiga', 'campaign-2-14', '2022-10-15 14:38:05', '2022-10-15 14:38:05');

-- --------------------------------------------------------

--
-- Table structure for table `campaign_images`
--

CREATE TABLE `campaign_images` (
  `id` int(50) NOT NULL,
  `campaign_id` int(50) DEFAULT NULL,
  `file_name` varchar(50) DEFAULT NULL,
  `is_primary` tinyint(50) DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `campaign_images`
--

INSERT INTO `campaign_images` (`id`, `campaign_id`, `file_name`, `is_primary`, `created_at`, `updated_at`) VALUES
(9, 9, 'images/13-bg.jpg', 1, '2022-10-15', '2022-10-15');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(50) UNSIGNED NOT NULL,
  `campaign_id` int(50) DEFAULT NULL,
  `user_id` int(50) DEFAULT NULL,
  `amount` int(50) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` date DEFAULT NULL,
  `payment_url` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `campaign_id`, `user_id`, `amount`, `status`, `code`, `created_at`, `updated_at`, `payment_url`) VALUES
(1, 9, 13, 100000, 'paid', '', '2022-10-13', '2022-10-13', NULL),
(2, 9, 13, 500000, 'paid', '', '2022-10-13', '2022-10-13', NULL),
(3, 10, 14, 500000, 'paid', '', '2022-10-13', '2022-10-13', NULL),
(4, 9, 13, 500000, 'pendding', '', '2022-10-16', '2022-10-16', NULL),
(5, 9, 13, 1000000, 'pendding', '', '2022-10-16', '2022-10-16', NULL),
(6, 9, 13, 123456789, 'pendding', '', '2022-10-18', '2022-10-18', 'https://app.sandbox.midtrans.com/snap/v3/redirection/22a3953e-eef0-44a9-8ab9-0c636d729472'),
(7, 9, 13, 101202, 'pendding', '', '2022-10-18', '2022-10-18', 'https://app.sandbox.midtrans.com/snap/v3/redirection/2373cce5-6210-4453-9a5c-a05090021d4e'),
(8, 9, 13, 10000, 'pendding', '', '2022-10-19', '2022-10-19', 'https://app.sandbox.midtrans.com/snap/v3/redirection/e2facb82-54fb-470b-9942-dd356dc280b6');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(50) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `occupation` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `password_hash` varchar(100) DEFAULT NULL,
  `avatar_file_name` varchar(100) DEFAULT NULL,
  `role` varchar(50) DEFAULT NULL,
  `token` varchar(100) DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `token`, `created_at`, `updated_at`) VALUES
(13, 'hendaru hery', 'hore', 'h@g.com', '$2a$04$phXxwaNmTxYeD43vfDhZMeLIx0XAQ42Cs06XkJr9LAbLDAB/PSVay', 'images/13-394873_3776810393454_474371649_n.jpg', 'user', NULL, '2022-10-15', '2022-10-15'),
(14, 'hendaru', 'hore', 'i@g.com', '$2a$04$RF6DthqRUB7cMaHbuwz81.wfVd55lZsLWb2Ls./V6JrpETAD35ege', '', 'user', NULL, '2022-10-15', '2022-10-15');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `campaign_images`
--
ALTER TABLE `campaign_images`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int(50) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `campaign_images`
--
ALTER TABLE `campaign_images`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(50) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
