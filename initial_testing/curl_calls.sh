

number_of_months_to_fetch=12

month_to_fetch=10
day_to_fetch=27
year_to_fetch=2025




for ((i=1; i<=number_of_months_to_fetch; i++)); do
    echo "$i"
    curl -s https://orthocal.info/api/gregorian/$year_to_fetch/$i/ > $i.json
    # curl -s https://orthocal.info/api/gregorian/$year_to_fetch/$month_to_fetch/$i/ | jq '.weekday'
    # curl -s https://orthocal.info/api/gregorian/$year_to_fetch/$month_to_fetch/$i/ | jq '.fast_level_desc'
    # curl -s https://orthocal.info/api/gregorian/$year_to_fetch/$month_to_fetch/$i/ | jq '.feast_level_description'
    # curl -s https://orthocal.info/api/gregorian/$year_to_fetch/$month_to_fetch/$i/ | jq '.fast_level'
    # curl -s https://orthocal.info/api/gregorian/$year_to_fetch/$month_to_fetch/$i/ | jq '.titles' | jq '.[]'
    echo "---"
done