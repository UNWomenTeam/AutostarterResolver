# AutostarterResolver - Testing tool
### Тест работы автозапуска
При старте появится файл или допишется, ели автозапуск сработал  

Ubuntu 20.04 its work  
Создадим init скрипт для запуска:    
`sudo vim /etc/init.d/autostarterResolver`  
Даем права на запуск (на исполнение):  
`sudo chmod 755 /etc/init.d/autostarterResolver`  
Запускаем:  
update-rc.d autostarterResolver defaults  

