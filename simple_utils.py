# simple_utils.py - A tiny utility library
def reverse_string(text):
    """
    Return the input string with characters in reverse order.
    
    Parameters:
        text (str): The string to reverse.
    
    Returns:
        reversed_text (str): The reversed string.
    """
    return text[::-1]
def count_words(sentence):
    """
    Count the words in a sentence.
    
    Words are determined by splitting the input string on whitespace.
    
    Parameters:
        sentence (str): The text in which to count words.
    
    Returns:
        int: The number of words in `sentence`.
    """
    return len(sentence.split())
def celsius_to_fahrenheit(celsius):
    """
    Convert a Celsius temperature to Fahrenheit.
    
    Parameters:
        celsius (float | int): Temperature in degrees Celsius.
    
    Returns:
        float: Temperature in degrees Fahrenheit corresponding to the input.
    """
    return (celsius * 9/5) + 32