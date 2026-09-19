# LeetCode's Official Editorial

### Approach 1: Divide into Regions

#### Intuition

To determine whether the circle and rectangle overlap, consider the boundary case where they share exactly one point. There are two possibilities:

- The circle is tangent to one of the four sides of the rectangle.
- The circle touches one of the rectangle's vertices.

Let the center of the circle be (x,y), its radius be r, the lower-left vertex of the rectangle be (x1​,y1​), and its upper-right vertex be (x2​,y2​). Formally, the two cases can be subdivided as follows:

- The circle is tangent to one of the four sides of the rectangle:
 - x\=x1​−r, y∈\[y1​,y2​\].
 - x\=x2​+r, y∈\[y1​,y2​\].
 - y\=y1​−r, x∈\[x1​,x2​\].
 - y\=y2​+r, x∈\[x1​,x2​\].
- The circle touches one of the rectangle's vertices:
 - (x−x1​)2+(y−y1​)2\=r2, x∈\[x1​−r,x1​\], y∈\[y1​−r,y1​\].
 - (x−x2​)2+(y−y1​)2\=r2, x∈\[x2​,x2​+r\], y∈\[y1​−r,y1​\].
 - (x−x1​)2+(y−y2​)2\=r2, x∈\[x1​−r,x1​\], y∈\[y2​,y2​+r\].
 - (x−x2​)2+(y−y2​)2\=r2, x∈\[x2​,x2​+r\], y∈\[y2​,y2​+r\].

Thus, the boundary of the valid positions of the circle's center forms a rounded rectangle. If the circle's center lies inside or on the boundary of this rounded rectangle, a common point exists; if it lies outside, no common point exists. The problem is therefore transformed into determining whether the circle's center lies inside or on this rounded rectangle.

We can divide the rounded rectangle into nine regions:

- Center rectangle: x∈\[x1​,x2​\], y∈\[y1​,y2​\].
- Upper rectangle: x∈\[x1​,x2​\], y∈\[y2​,y2​+r\].
- Lower rectangle: x∈\[x1​,x2​\], y∈\[y1​−r,y1​\].
- Left rectangle: x∈\[x1​−r,x1​\], y∈\[y1​,y2​\].
- Right rectangle: x∈\[x2​,x2​+r\], y∈\[y1​,y2​\].
- Lower-left rounded corner: (x−x1​)2+(y−y1​)2≤r2, x∈\[x1​−r,x1​\], y∈\[y1​−r,y1​\].
- Lower-right rounded corner: (x−x2​)2+(y−y1​)2≤r2, x∈\[x2​,x2​+r\], y∈\[y1​−r,y1​\].
- Upper-left rounded corner: (x−x1​)2+(y−y2​)2≤r2, x∈\[x1​−r,x1​\], y∈\[y2​,y2​+r\].
- Upper-right rounded corner: (x−x2​)2+(y−y2​)2≤r2, x∈\[x2​,x2​+r\], y∈\[y2​,y2​+r\].

We examine these cases separately. When checking the rounded corners, since the five rectangular cases have already been checked, we do not need to separately examine the ranges of the horizontal and vertical coordinates; it is sufficient to directly compare whether the squared distance from the circle's center to the corresponding vertex is at most r2.

#### Complexity Analysis

- Time complexity: O(1).

- Space complexity: O(1).

### Approach 2: Minimum Distance from the Circle's Center to the Rectangle

#### Intuition

When determining the positional relationship between a circle and a line, we often compute the length of the perpendicular segment from the circle's center to the line. If this distance is less than or equal to the radius, the circle and the line intersect. More precisely, this works because the perpendicular segment is the shortest segment from the circle's center to any point on the line. If this minimum length is at most the circle's radius, then at least one point on the line lies on or inside the circle.

We can apply the same idea to compute the minimum distance from the circle's center to the rectangular region.

Split this into two problems: the minimum distance xmin​ from the circle's center to the region x1​≤x≤x2​, and the minimum distance ymin​ from the circle's center to the region y1​≤y≤y2​. We obtain

xmin​(x)\={0,min{∣x−x1​∣,∣x−x2​∣},​x∈\[x1​,x2​\],otherwise.​

ymin​(y)\={0,min{∣y−y1​∣,∣y−y2​∣},​y∈\[y1​,y2​\],otherwise.​

The minimum distance from the circle's center to the rectangular region is xmin2​+ymin2​​. For ease of understanding, the plane can be divided into nine regions by the four lines x\=x1​, x\=x2​, y\=y1​, and y\=y2​. Discussing the cases separately yields this combined result.

After obtaining the distance, compare it with the radius. If the distance is no greater than the radius, a common point exists.

#### Complexity Analysis

- Time complexity: O(1).

- Space complexity: O(1).

---
