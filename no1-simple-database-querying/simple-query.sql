/*
Soal : 
——————————————————————————
| ID | UserName | Parent |
——————————————————————————
| 1 | Ali | 2 |
| 2 | Budi | 0 |
| 3 | Cecep | 1 |
—————————————————————————-
Tuliskan SQL Query untuk mendapatkan data berisi:
——————————————————————————————————
| ID | UserName | ParentUserName |
——————————————————————————————————
| 1 | Ali | Budi |
| 2 | Budi | NULL |
| 3 | Cecep | Ali |
——————————————————————————————————


Jawaban :
Asumsi Nama Table user: 
*/
CREATE TABLE `user`  (
  `ID` int(0) NOT NULL,
  `UserName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Parent` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

/*Select Query : */

select u.ID, u.UserName, p.UserName as ParentUserName 
               from user as u left join user as p on u.Parent = p.ID

/*Result
| 1 | Ali | Budi |
| 2 | Budi | NULL |
| 3 | Cecep | Ali |*/

