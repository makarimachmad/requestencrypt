-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 05, 2021 at 12:51 AM
-- Server version: 10.4.19-MariaDB
-- PHP Version: 8.0.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `webinar`
--

-- --------------------------------------------------------

--
-- Table structure for table `database_keys`
--

CREATE TABLE `database_keys` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `username` varchar(25) NOT NULL,
  `password` varchar(255) NOT NULL,
  `encryption_key` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `database_keys`
--

INSERT INTO `database_keys` (`id`, `user_id`, `username`, `password`, `encryption_key`) VALUES
(3, 7, 'pengunjung1', 'caf1a3dfb505ffed0d024130f58c5cfa', '123');

-- --------------------------------------------------------

--
-- Table structure for table `pengunjungs`
--

CREATE TABLE `pengunjungs` (
  `id` int(11) NOT NULL,
  `nama` varchar(30) NOT NULL,
  `umur` int(11) NOT NULL,
  `alamat` varchar(150) NOT NULL,
  `pekerjaan` varchar(25) NOT NULL,
  `tanggal_lahir` date NOT NULL,
  `created_date` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `pengunjungs`
--

INSERT INTO `pengunjungs` (`id`, `nama`, `umur`, `alamat`, `pekerjaan`, `tanggal_lahir`, `created_date`) VALUES
(1, 'tes', 23, 'jogja', 'se', '1998-08-02', '2021-11-03 07:08:39'),
(3, 'gitu ya', 48, 'surabaya', 'dsa', '1973-09-28', '2021-11-03 07:32:05'),
(4, 'satudua', 48, 'malang', '', '1973-09-28', '2021-11-03 08:42:39'),
(5, 'apa', 48, 'dimana', 'guru', '1973-09-28', '2021-11-04 06:49:49'),
(6, 'apa', 48, 'dimana', 'guru', '1973-09-28', '2021-11-04 13:34:47'),
(7, 'apa', 48, 'dimana', 'guru', '1973-09-28', '2021-11-04 13:45:40');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `nama` varchar(50) NOT NULL,
  `alamat` varchar(255) NOT NULL,
  `no_hp` varchar(13) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `nama`, `alamat`, `no_hp`) VALUES
(1, 'sasa', 'pekalongan', '11221122'),
(5, 'messi', 'paris', '883321'),
(12, 'bukan mbappe', 'madrid', '777777'),
(13, 'bapak saya', 'pandeyan', '8483838');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `database_keys`
--
ALTER TABLE `database_keys`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `pengunjungs`
--
ALTER TABLE `pengunjungs`
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
-- AUTO_INCREMENT for table `database_keys`
--
ALTER TABLE `database_keys`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `pengunjungs`
--
ALTER TABLE `pengunjungs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `database_keys`
--
ALTER TABLE `database_keys`
  ADD CONSTRAINT `database_keys_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `pengunjungs` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
