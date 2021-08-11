# SEI调试指令

播放视频源

adb shell

su

am start -n com.utsmta.app/com.nes.factorytest.ui.activity.MediaActivity

input keyevent 66



boxid adb 写入

cd /sys/class/unifykeys/

cat list

echo sboxid > name

cat read

echo sboxid > name

echo 785965489654 > write

cat read