// http://repl.it/oTp/18
// http://www.careercup.com/question?id=6303093824159744
import java.lang.*;
import java.util.*;

class Main {
    
    public static void main(String[] args) {
        List<Kid> kids = new ArrayList<>();
        for (int i = 49; i < 99; i++) {
            kids.add(new Kid("John Doe " + (i + 1), i + 1));
        }
        
        for (int i = 49; i < 99; i++) {
            kids.add(new Kid("John Doe " + (i + 1), i + 1));
        }

        Collections.shuffle(kids, new Random());
        
        System.out.println(canoesCount(kids) + " canoes required.");
    }
    
    
    public static int canoesCount(List<Kid> kids) {
        Collections.sort(kids, new KidComparator());
        
        int front = 0;
        int back = kids.size() - 1;
        int canoeCount = 0;
        
        while (kids.get(back).getWeight() > 150) {
            // Trim the fat kids.
            back--;
        }
        
        while (back >= front) {
            canoeCount++;
            
            if (kids.get(front).getWeight() + kids.get(back).getWeight() <= 150) {
                front++;
            }

            back--;
        }        
        
        return canoeCount;
    }
}


class KidComparator implements Comparator<Kid> {

    public int compare(Kid o1, Kid o2) {
        return o1.getWeight() - o2.getWeight();
    }    
}


class Kid {
    
    private String name;
    private int weight;
    
    
    public Kid(String name, int weight) {
        this.name = name;
        this.weight = weight;
    }
    
    
    public String toString() {
        return "[Name: " + name + ", Weight: " + weight + "]";    
    }
    
    
    public String getName() {
        return name;
    }
    
    
    public int getWeight() {
        return weight;
    }
}
