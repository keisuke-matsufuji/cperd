REVIEWDOG_COMMAND="reviewdog -efm=\"%f:%l:%c: %m\" -diff=\"git diff ${GITHUB_REF}\""

TARGET_DIR="./"
for file in $(find $TARGET_DIR -name "php.ini" -o -name ".htaccess" -o -name "*.php")
do
  # 各php.iniファイルからerror_reportingの値を取得
  ERROR_REPORTING=$(grep error_reporting $file | cut -d= -f2 | tr -d ' ')
  echo $ERROR_REPORTING

  # その値をmain.goに渡し、出力を取得
  OUTPUT=$(go run main.go $ERROR_REPORTING)
  echo $OUTPUT

  # その出力をreviewdogに渡す
  echo $OUTPUT | $REVIEWDOG_COMMAND
done

