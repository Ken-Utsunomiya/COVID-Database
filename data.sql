INSERT INTO Hospital
VALUES
(1,'899 W 12th Ave, Vancouver, BCV5Z 1M9','Saint Paul’s Hospital','6047863712', 50),
(2,'1081 Burrard St, Vancouver, BCV6Z 1Y6','Vancouver General Hospital','4187911084', 50),
(3,'1004 West Blvd, Vancouver, BCV6G 1W6','UBC Student Health Clinic','6043956928', 75),
(4,'3080 Prince Edward St, Vancou-ver, BC V5T 3N4','Saint Jones Hospital','6933853958', 100),
(5, '4500 Oak St, Vancouver, BC V6H3N1', 'Vancouver Children’s Hospital', '5029535822', 30);

INSERT INTO Bed_Contain
VALUES
('H0001','B1011','BU01'),
('H0001','B1012','BU01'),
('H0001','B2301','BU01'),
('H0002','B1021','BU03'),
('H0002','B3031','BU03'),
('H0003','B5381','BU02'),
('H0004', 'B2301', 'BU01');

INSERT INTO Staff_Employee
VALUES
('H0001', 'D0003', 'Janes Ng', '3485923409', '1969-06-16'),
('H0001', 'G0001', 'Mandy Asia', NULL, '1996-02-21'),
('H0001','D0001','Tamirrice Black', '4187911095', '1996-02-21'),
('H0001','N0001','Freddie Lives','4187911084','1996-02-21'),
('H0002','D0025','Shelly Matter','6043956928','1989-10-12'),
('H0002','N0123', 'Wally Kim', '3066345678', '1968-06-11'),
('H0003','D0030','George Floyd','6933853958','1996-02-21'),
('H0004','D0002','Terence Crutcher','5029535822','1996-02-21'),
('H0003','N0029', 'Emily Till', '3925839203', '1999-03-13'),
('H0004','N0080', 'Eric Garner', '6047863811', '1996-02-21'),
('H0005','N0110', 'Ahmaud Arbery', '0611202001', '1980-06-22'),
('H0005', 'D0193', 'James Wong', NULL, '1960-10-07'),
('H0001','N0002', 'James Wong', NULL, '1955-10-06'),
('H0002','N0124', 'India Kager', '0394230948', '1998-08-31'),
('H0003','N0030', 'Miguel Espinal', '4903410879', '1996-02-29'),
('H0004','N0081', 'Walter Smart', '3059720233', '1998-02-19'),
('H0005','N0111', 'Billy Ray Davis', '8032470234', '1990-08-29');

INSERT INTO Doctor_Employee
VALUES
('H0001','D0001','Pediatrician'),
('H0001','D0003','Emergency medicine specialist'),
('H0002','D0025','Pathologist'),
('H0003','D0030','Internist'),
('H0004','D0002','Infectious disease specialist'),
('H0005','D0193','Pediatrician');

INSERT INTO Nurse_Employee
VALUES
('H0001','N0001'),
('H0002','N0123'),
('H0003','N0029'),
('H0004','N0080'),
('H0005','N0110'),
('H0001','N0002'),
('H0002','N0124'),
('H0003','N0030'),
('H0004','N0081'),
('H0005','N0111');

INSERT INTO Nurse_Supervise
VALUES
('H0001','N0001','N0002'),
('H0002','N0123','N0124'),
('H0003','N0029','N0030'),
('H0004','N0080', NULL),
('H0005','N0110', NULL);

INSERT INTO Health_Record
VALUES
('R0001', 'prozac', NULL),
('R0002','remdesivir','peanut'),
('R0030','remdesivir', NULL),
('R0308','prozac','shrimp'),
('R0500','hydroxychloroquine','dairy');

INSERT INTO Patient
VALUES
('P0001','H0001','D0001','BU01','B1011','R0001', TRUE,'John Snow','1980-05-31',FALSE, 172.1 ,65.2,'1209 Lower Mall','2020-04-01','2020-03-31'),
('P0002','H0001','D0001','BU01','B1012','R0002', TRUE,'Cersei Lannister','1990-2-27',TRUE, 165.0, 60.1,'201 W 10th Ave','2020-04-01','2020-03-30'),
('P0003','H0004','D0002','BU01','B2301', 'R0030',FALSE,'Khal Drogo','1975-12-11',FALSE, 180.5, 80.5,'2301 West Mall','2020-04-23','2020-04-20'),
('P0004','H0002','D0025','BU03','B1021','R0500', TRUE,'Sansa Stark','1985-09-20',TRUE, 163.4, 59.3,'12 Main Mall','2020-03-13','2020-03-13'),
('P0005','H0003','D0030','BU02','B5381', 'R0308', FALSE,'Tyrion Lannister','1990-07-31',FALSE, 178.1, 75.2,'13 Agronomy Rd','2020-03-28','2020-03-25');

INSERT INTO Guardian
VALUES
('6024939294','Eric Garner','father'),
('2059496883','John Crawford III','father'),
('1049482945','Michael Brown','grand father'),
('5062946020','Ezell Ford','cousin'),
('4023931850','Christian Taylor','brother');

INSERT INTO Coverage
VALUES
('II001',5000.0),
('II002',7000.0),
('II003',10000.0),
('II004', 8000.0),
('II193',75000.0);

INSERT INTO Insurance (patient_ID, insurance_ID, expiration_date)
VALUES
('P0001','II001','2022-12-31'),
('P0001','II002','2021-01-31'),
('P0002','II002','2020-10-23'),
('P0003','II001','2021-01-31'),
('P0005','II193','2021-01-31'),
('P0005','II004','2021-01-31');

INSERT INTO Shift
VALUES
('00:00',5),
('05:00',6),
('11:00',6),
('17:00',7),
('18:00',6);

INSERT INTO Assign
VALUES
('P0001','H0001','N0001','05:00'),
('P0002','H0001','N0002','05:00'),
('P0003','H0004','N0080','11:00'),
('P0004','H0002','N0123','17:00'),
('P0005','H0003','N0029','00:00');

INSERT INTO Take_Care
VALUES
('P0001','6024939294','Eric Garner'),
('P0002','2059496883','John Crawford III'),
('P0002','1049482945','Michael Brown'),
('P0003','5062946020','Ezell Ford'),
('P0005','4023931850','Christian Taylor');

INSERT INTO Test_Lab
VALUES
('L0001A','PCR'),
('L0001C','Antibody test'),
('L0003B','PCR'),
('L0004A','PCR'),
('L0005B','Antibody test');

INSERT INTO Test_Order
VALUES
('T0001','P0001','H0001','D0001',FALSE,'2020-03-15','L0001C'),
('T0002','P0001','H0001','D0001',FALSE,'2020-03-16','L0003B'),
('T0003','P0002','H0001','D0001',FALSE,'2020-04-01','L0001C'),
('T0004','P0005','H0003','D0030',TRUE,'2020-03-27','L0003B'),
('T0005','P0003','H0004','D0002',TRUE,'2020-04-20','L0005B'),
('T0006','P0004','H0002','D0025',FALSE,'2040-04-22','L0001A');