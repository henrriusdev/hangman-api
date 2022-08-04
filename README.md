# HangMan API

This API is **easy to use**, since it is designed for data consumption, although you can also enter hints with their respective responses (you will see this in depth later). The API is in charge of **returning you the hint together with its answer and the length of this answer** thanks to the fact that **you will pass it a value through a *URL***, the value would be the ID of the record to which you want to obtain the values mentioned above.

The API website is: https://hangmanapi-hb.herokuapp.com/

## Get Hangs

Getting your '*Hang*' (I'll call the **returned hint, response and response length**) is very simple, just use the following URL template:

    https://hangmanapi-hb.herokuapp.com/get/hangman?id=X
   
Remember that you must change the X for the ID that you want. Now, let's try to receive the **data from the API**, for this, I will use JS, first create an async function, inside this make a try-catch block and inside the try block declare a constant with the name "result", that will store the promise of the fetch to the API website, and after making this line of code, declare a constant with the name you want to store the value of the result.json(). If you are confused, take a guide from the following code:

    async function getHangman()  {
	    try {
			const result = await fetch(
				"https://hangmanapi-hb.herokuapp.com/get/hangman?id=1",{ method:  "GET",}
			);
			const  db  =  await result.json();
			console.log(db);
		}  catch (e) {
			console.log(e);
		}
	}
	
	getHangman();

And then, **open the console of your web browser** and you will notice that you have the JSON of your '*Hang*', then from here, it's up to you how you design the game. But to help you, I will tell you the specific function of each variable.

 - Hint: store the hint so you can guess and there is no hangman, note that it MUST be in English, you could buy a coffee if you want me to store hintin other languages ;)
 - Answer: stores the answer to the given hint, this SHOULD also be in English.
 - AnswerLenght: stores the length of the response, obviously. This will help you create the necessary spaces to display to the user.

## Create Hangs
To create a Hang, the model to follow is very similar, you only have to change the ID parameter for the hint and answer parameters, as you can see in the following example:

    https://hangmanapi-hb.herokuapp.com/new/hangman?hint=new+hint&answer=hint+answer
   
 Note that if you want the parameters to have spaces, you must **concatenate them using the plus symbol**, after you have finished passing a parameter, continue using "&", again I remind you that if you want this API to control more languages, you can buy me a coffee ;)

