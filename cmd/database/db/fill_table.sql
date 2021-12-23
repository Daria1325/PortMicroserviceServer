-- Set params
set session my.name = 'consequat id officia';
set session my.is_active = false;
set session my.company = 'GEOFARM';
set session my.email = 'maymclean@geofarm.com';
set session my.phone = '+1 (997) 434-3843';
set session my.about = 'Id laborum labore irure nisi mollit. Exercitation dolor ad nisi veniam tempor laboris Lorem nisi incididunt do reprehenderit veniam dolor consequat. Mollit deserunt occaecat tempor fugiat consequat culpa eu eu deserunt minim qui. Dolore magna ipsum nisi est occaecat deserunt aliquip ';
set session my.registered = '2021-02-14T01:46:57 -02:00';
set session my.latitude = 70.822864;
set session my.longitude = 156.088083;
set session my.address = '907 National Drive, Foscoe, Oregon, 5061';


-- Filling of products
INSERT INTO ports (name, is_active,company,email,phone,address,about,registered,latitude,longitude)
VALUES ( 'my.name' , 'my.is_active', 'my.company', 'my.email', 'my.phone', 'my.address', 'my.about', 'my.registered', 'my.latitude', 'my.longitude');
