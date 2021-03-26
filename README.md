# Ascii-Art-Web
It's a web server for outputting the text in a graphic representation of ASCII and dowloading a file with it.
# execute: go run main.go

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of a program that outputs text or a picture in a graphical representation of ASCII

Our site translate pictures and text into ASCII caracters. 

## User Manual

- You have to run the program in the terminal.

go run main.go

- The program will run in local host, you have to take the adress of the host and paste it in your browser.

- You enter a sentence, a letter, a number or a picture onto the site to see it translated as it is in exemples.

- You can choose if you want to print it on a text file.

- You can also chose the name file.

- Then you submit.

## Here how the site is presented

![picture](https://git.ytrack.learn.ynov.com/ACLAVERIA/ascii-art-web/raw/branch/master/ScreenReadme/Welcome.png)

### Exemples

----------------------------------------------------------

1. **No type on the textarea**

It just print "Type Something".

![picture](https://git.ytrack.learn.ynov.com/ACLAVERIA/ascii-art-web/raw/branch/master/ScreenReadme/First.png)

---------------------------------------------------------

2. **With a word but not first test not deleted and inside a file**

It print "Type something + word".

Exemple with "Hello" and the file name is TypeText

![picture](https://git.ytrack.learn.ynov.com/ACLAVERIA/ascii-art-web/raw/branch/master/ScreenReadme/TypeSomething.png)
![picture](https://git.ytrack.learn.ynov.com/ACLAVERIA/ascii-art-web/raw/branch/master/ScreenReadme/InFile.png)

---------------------------------------------------------

3. **Just a word and a different font style. Not printed on a file**

Only the word will be printed.

Exemple with "Debug" and "shadow"

![picture](https://git.ytrack.learn.ynov.com/ACLAVERIA/ascii-art-web/raw/branch/master/ScreenReadme/debugsha.png)
---------------------------------------------------------

4. **Test with word + number + special character + font style and in file**

It will simply add the Ascii art of the number to the Ascii art of the translated word.

We'll try with "Test 4 $" in a font style Thinkertoy and in a file named Test4

![picture](https://git.ytrack.learn.ynov.com/ACLAVERIA/ascii-art-web/raw/branch/master/ScreenReadme/Test4.png)
![picture](https://git.ytrack.learn.ynov.com/ACLAVERIA/ascii-art-web/raw/branch/master/ScreenReadme/text.png)

---------------------------------------------------------

5. **With a picture** 

The picture will be translated into Ascii art and will be print on a new page of the browser.

We choose the option "Phoque" which mean seal.

![picture](https://git.ytrack.learn.ynov.com/ACLAVERIA/ascii-art-web/raw/branch/master/ScreenReadme/Seal.png)

---------------------------------------------------------

6. **With more than one line**

For this time we tried with more lines. It works for 2 but it's the same for 3 or more.

The exemple contain the \n : "Hello\nThere"

![picture](https://git.ytrack.learn.ynov.com/ACLAVERIA/ascii-art-web/raw/branch/master/ScreenReadme/Hello.png)

---------------------------------------------------------

7. **If Index out of range***

If you want to type something outside the ASCII table or if you make a backslash n on the textarea

![picture](https://git.ytrack.learn.ynov.com/ACLAVERIA/ascii-art-web/raw/branch/master/ScreenReadme/index.png)
