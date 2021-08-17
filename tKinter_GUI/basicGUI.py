from tkinter import *

root = Tk()

# Input field
e = Entry(root)
e.pack()

# Create label widget
#myLabel = Label1(root, text="Hello World!")
#myLabel = Label2(root, text="BlahBleeBloo")

# Populate grid
#myLabel1.grid(row=0, column=0)
#myLabel2.grid(row=1, column=0)

def myClick():
    myLabel = Label(root, text="Meepins, " + e.get())
    myLabel.pack()

# Create button, use padx= and pady= to resize button, state= to enable/dis button
myButton = Button(root, text="Click Me.....Plz?", command = myClick)
myButton.pack()

root.mainloop()