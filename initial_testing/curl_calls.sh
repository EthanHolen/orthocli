

number_of_entries_to_fetch=30

month_to_fetch=12
day_to_fetch=09
year_to_fetch=2024




for ((i=1; i<=number_of_entries_to_fetch; i++)); do
    # echo "$i"
    curl -s https://orthocal.info/api/gregorian/$year_to_fetch/$month_to_fetch/$i/ | jq '.weekday'
    curl -s https://orthocal.info/api/gregorian/$year_to_fetch/$month_to_fetch/$i/ | jq '.titles' | jq '.[]'
done