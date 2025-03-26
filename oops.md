Certainly! Let's go over the **OOP concepts** using **C# (.NET)**, with clear definitions and examples for each concept.

---

### **1. Encapsulation**

**Definition**: Encapsulation is the process of combining data (variables) and methods (functions) into a single unit, called a class. It also involves restricting access to some of the object’s components to protect the integrity of the object’s state by using access modifiers like `private`, `protected`, and `public`.

**Example**:
```csharp
using System;

public class Person
{
    // Private fields
    private string name;
    private int age;

    // Public method to access private fields
    public string GetName()
    {
        return name;
    }

    public void SetName(string n)
    {
        name = n;
    }

    public int GetAge()
    {
        return age;
    }

    public void SetAge(int a)
    {
        if (a > 0) age = a;  // Simple validation for age
    }
}

class Program
{
    static void Main(string[] args)
    {
        Person person = new Person();
        person.SetName("John");
        person.SetAge(30);

        Console.WriteLine("Name: " + person.GetName());  // Output: Name: John
        Console.WriteLine("Age: " + person.GetAge());    // Output: Age: 30
    }
}
```
- **Encapsulation**: The fields `name` and `age` are private, and we use public methods (`SetName`, `GetName`, etc.) to access and modify them.

---

### **2. Inheritance**

**Definition**: Inheritance is a mechanism where one class (child class) can inherit the properties and methods of another class (parent class). This allows for code reuse and extension of functionality.

**Example**:
```csharp
using System;

public class Animal
{
    public void Sound()
    {
        Console.WriteLine("Animal makes a sound");
    }
}

public class Dog : Animal   // Inherits from Animal class
{
    public void Sound()     // Overriding the Sound method
    {
        Console.WriteLine("Dog barks");
    }
}

class Program
{
    static void Main(string[] args)
    {
        Dog dog = new Dog();
        dog.Sound();  // Output: Dog barks
    }
}
```
- **Inheritance**: The `Dog` class inherits from `Animal`. It overrides the `Sound()` method to provide its own implementation.

---

### **3. Polymorphism**

**Definition**: Polymorphism allows objects of different classes to be treated as objects of a common superclass. The actual method that gets called is determined at runtime, allowing for different implementations of the same method.

**Example**:
```csharp
using System;

public class Animal
{
    public virtual void Sound()  // Virtual method to allow overriding
    {
        Console.WriteLine("Animal makes a sound");
    }
}

public class Dog : Animal
{
    public override void Sound()  // Method overriding
    {
        Console.WriteLine("Dog barks");
    }
}

public class Cat : Animal
{
    public override void Sound()  // Method overriding
    {
        Console.WriteLine("Cat meows");
    }
}

class Program
{
    static void Main(string[] args)
    {
        Animal animal;

        animal = new Dog();
        animal.Sound();  // Output: Dog barks

        animal = new Cat();
        animal.Sound();  // Output: Cat meows
    }
}
```
- **Polymorphism**: Both `Dog` and `Cat` override the `Sound()` method from `Animal`. The method called is determined by the actual object type at runtime, allowing different behaviors.

---

### **4. Abstraction**

**Definition**: Abstraction is the concept of hiding complex implementation details and showing only the essential features to the user. In C#, this is achieved by using abstract classes and interfaces.

**Example**:
```csharp
using System;

public abstract class Shape  // Abstract class
{
    public abstract void Draw();  // Abstract method (no implementation)
}

public class Circle : Shape
{
    public override void Draw()  // Providing implementation
    {
        Console.WriteLine("Drawing Circle");
    }
}

public class Square : Shape
{
    public override void Draw()  // Providing implementation
    {
        Console.WriteLine("Drawing Square");
    }
}

class Program
{
    static void Main(string[] args)
    {
        Shape shape;

        shape = new Circle();
        shape.Draw();  // Output: Drawing Circle

        shape = new Square();
        shape.Draw();  // Output: Drawing Square
    }
}
```
- **Abstraction**: The `Shape` class is abstract, meaning it cannot be instantiated directly. The `Circle` and `Square` classes provide concrete implementations of the `Draw()` method.

---

### **5. Composition**

**Definition**: Composition is a design principle where a class is composed of one or more objects of other classes. This "has-a" relationship is used to build complex behavior by reusing objects.

**Example**:
```csharp
using System;

public class Engine
{
    public void Start()
    {
        Console.WriteLine("Engine starting...");
    }
}

public class Car
{
    private Engine engine;  // Composition: Car has an Engine

    public Car()
    {
        engine = new Engine();  // Initialize engine
    }

    public void Start()
    {
        engine.Start();  // Delegation: Car delegates the start functionality to Engine
        Console.WriteLine("Car starting...");
    }
}

class Program
{
    static void Main(string[] args)
    {
        Car car = new Car();
        car.Start();  // Output: Engine starting... Car starting...
    }
}
```
- **Composition**: The `Car` class has an `Engine` object. The `Car` class delegates the start functionality to the `Engine` class.

---

### Summary of OOP Concepts in C#:

| Concept        | Definition                                             | C# Example                                         |
|----------------|--------------------------------------------------------|-----------------------------------------------------|
| **Encapsulation** | Wrapping data (fields) and methods into a single unit (class), and controlling access. | Private fields with public getter/setter methods.  |
| **Inheritance**  | A class inherits properties and methods from another class. | `class Dog : Animal` for inheritance.              |
| **Polymorphism** | Methods can have different implementations based on the object type at runtime. | Method overriding using `virtual` and `override`.  |
| **Abstraction**  | Hiding the complex implementation and exposing only necessary functionality. | Abstract classes with `abstract` methods.         |
| **Composition**  | A class is made up of other objects to build more complex functionality. | `Car` class has an `Engine` object.                |

---

### **Key Differences with C# OOP Features**:
- **Classes and Objects**: C# uses **classes** to define the blueprint and **objects** as instances of the class.
- **Virtual and Abstract Methods**: In C#, polymorphism is achieved through **virtual** (for base class methods) and **abstract** methods (for methods that must be implemented by derived classes).
- **Access Modifiers**: C# supports access modifiers like `private`, `protected`, and `public` to control the visibility of class members.

C# provides a rich set of OOP features, which allow developers to write highly modular, maintainable, and reusable code by adhering to OOP principles such as **encapsulation**, **inheritance**, **polymorphism**, **abstraction**, and **composition**.