package iastronautpairs

// Astronaut Pairs
// The member states of the UN are planning to send people to the moon.
// They want them to be from different countries. You will be given a list of pairs of astronaut ID's.
//  Each pair is made of astronauts from the same country.
// Determine how many pairs of astronauts from different countries they can choose from.

// Example

// n = 4

// astronauts = [1,2] [2,3]

// There are 4 astronauts numbered 0 through 3. Astonauts grouped by country are [0] and [1,2,3].
// There are 3 pairs to choose from [0,1], [0,2] and [0,3].

// Input

// Input contains the number N : denoting number of astronauts and and list L denoting the pairs of astronauts from the same country.

// N = 5
// vector<pair<int,int> > astronauts = {{0,1}, {2,3} , {0,4}};
// Output

// 6

// There are 6 ways to choose a pair (0,2) (0,3) (1,2) (1,3) (4,2) and (4,3) as astronauts (0,1,4) belong to country 1 and (2,3) belong to another.
