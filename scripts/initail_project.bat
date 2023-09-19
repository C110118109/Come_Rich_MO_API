curl -fLo ../air.exe https://github.com/cosmtrek/air/releases/download/v1.43.0/air_1.43.0_windows_amd64.exe
curl -fLo migrate.zip https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.windows-amd64.zip
echo F | xcopy air.example.windows .air.toml /H
powershell Expand-Archive migrate.zip -DestinationPath ./
del migrate.zip
move .air.toml ../
move migrate.exe ../