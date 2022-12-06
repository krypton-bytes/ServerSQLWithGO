-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1:3306
-- Tiempo de generación: 05-12-2022 a las 08:43:48
-- Versión del servidor: 5.7.36
-- Versión de PHP: 7.4.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `restaurantego`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `platillos`
--

DROP TABLE IF EXISTS `platillos`;
CREATE TABLE IF NOT EXISTS `platillos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(2500) NOT NULL,
  `descripcion` varchar(2500) NOT NULL,
  `precio` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=13 DEFAULT CHARSET=latin1;

--
-- Volcado de datos para la tabla `platillos`
--

INSERT INTO `platillos` (`id`, `nombre`, `descripcion`, `precio`) VALUES
(1, 'Hamburguesa hawaiana', 'Emparedado', 90),
(4, 'Papas a la francesa', 'Papas picadas y hervidas en aceite que se sirven con ketchup y queso derretido', 50),
(5, 'Papas a la francesa', 'Papas picadas y hervidas en aceite que se sirven con ketchup y queso derretido', 50),
(6, 'Papas a la francesa', 'Papas picadas y hervidas en aceite que se sirven con ketchup y queso derretido', 50),
(7, 'Papas a la francesa', 'Papas picadas y hervidas en aceite que se sirven con ketchup y queso derretido', 50),
(8, 'Papas a la francesa', 'Papas picadas y hervidas en aceite que se sirven con ketchup y queso derretido', 50),
(9, 'Papas a la francesa', 'Papas picadas y hervidas en aceite que se sirven con ketchup y queso derretido', 50),
(10, 'Helado', 'Crema de nieve en un barquillo de galleta', 20),
(11, 'Papas a la francesa', 'Papas picadas y hervidas en aceite que se sirven con ketchup y queso derretido', 50);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
