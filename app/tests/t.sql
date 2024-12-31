--------------------------------------- SYSTEM TABLE ---------------------------------------

CREATE TABLE tb_products (
    product_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'รหัสสินค้า (Primary Key)', 
    prodmgr_id INT NULL COMMENT 'ผู้จัดการฝ่ายผลิต [mst_params_dtl.param_id=11]',      
    sku VARCHAR(20) NULL UNIQUE COMMENT 'รหัสสินค้า (SKU) ที่ไม่ซ้ำกัน',                     
    name VARCHAR(100) NULL COMMENT 'ชื่อสินค้า', 
    name_search VARCHAR(100) NULL COMMENT 'ชื่อในการค้นหา',  
    name_th VARCHAR(100) NULL COMMENT 'ชื่อสินค้าภาษาไทย',
    name_eng VARCHAR(100) NULL COMMENT 'ชื่อสินค้าภาษาอังกฤษ',                          
    category1 INT NULL COMMENT '[mst_params_dtl.param_id=1]',
    category2 INT NULL COMMENT '[mst_params_dtl.param_id=2]',                       
    category3 INT NULL COMMENT '[mst_params_dtl.param_id=3]',                       
    category4 INT NULL COMMENT '[mst_params_dtl.param_id=4]',                       
    unit_price DECIMAL(10, 2) NULL COMMENT 'ราคาต่อหน่วย (บาท)',   
    lower_price DECIMAL(10, 2) NULL COMMENT 'ราคาขั้นต่ำ (บาท)',    
    promotion_price DECIMAL(10, 2) NULL COMMENT 'ราคาสินค้าราคาพิเศษ (หากมี)',         
    stock_quantity DECIMAL(10, 2) NULL COMMENT 'ปริมาณสินค้าคงเหลือ (กิโลกรัม)',               
    supplier_id INT NULL COMMENT '[mst_params_dtl.param_id=5]',                                 
    weight DECIMAL(10, 2) COMMENT 'น้ำหนักต่อหน่วย (กิโลกรัม)',                            
    unit VARCHAR(20) DEFAULT 'กิโลกรัม' COMMENT 'หน่วยของสินค้า (เช่น กิโลกรัม หรือ ชิ้น)',            
    description TEXT NULL COMMENT 'รายละเอียดเพิ่มเติมของสินค้า',                            
    discount DECIMAL(5, 2) NULL COMMENT 'ส่วนลด (เปอร์เซ็นต์หรือจำนวนเงิน)',                     
    expiry_date DATE NULL COMMENT 'วันที่หมดอายุ (สำหรับสินค้าที่มีวันหมดอายุ)',                            
    stock_alert INT NULL COMMENT 'จำนวนสินค้าที่เหลือน้อยกว่าระดับนี้จะได้รับการแจ้งเตือน',                             
    brand_id INT NULL COMMENT 'แบรนด์สินค้า [mst_params_dtl.param_id=6]',     
    color_id INT NULL COMMENT 'สีสินค้า [mst_params_dtl.param_id=7]',
    size_id INT NULL COMMENT 'ไชต์สินค้า [mst_params_dtl.param_id=8]',
    packsize_id INT NULL COMMENT 'ขนาดกล่องภายนอก [mst_params_dtl.param_id=12]',
    innerpacksize_id INT NULL COMMENT 'ขนาดกล่องภายใน [mst_params_dtl.param_id=13]',
    qtypack_id INT NULL COMMENT 'จำนวนกล่องภายนอก [mst_params_dtl.param_id=14]',
    qtyinnerpack_id INT NULL COMMENT 'จำนวนกล่องภายใน [mst_params_dtl.param_id=15]',
    collection_id INT NULL COMMENT 'collection ของสินค้า [mst_params_dtl.param_id=16]',
    factory_id INT NULL COMMENT 'โรงงานที่ผลิต [mst_params_dtl.param_id=9]',
    compsells_id INT NULL COMMENT 'บริษัทที่จัดจำหน่าย [mst_params_dtl.param_id=10]',
    tax_rate DECIMAL(5, 2) NULL COMMENT 'อัตราภาษีสินค้า',                     
    barcode VARCHAR(50) NULL COMMENT 'รหัสบาร์โค้ดสินค้า (หากมี)',
    minorder_qty INT NULL COMMENT 'จำนวนสินค้าขั้นต่ำในการสั่งซื้อ',
    is_active BOOLEAN DEFAULT TRUE COMMENT 'สถานะการใช้งาน (TRUE = ใช้งาน, FALSE = ไม่ใช้งาน)', 
    created_by VARCHAR(100) NULL COMMENT 'ผู้สร้างข้อมูล', 
    updated_by VARCHAR(100) NULL COMMENT 'ผู้แก้ไขข้อมูลล่าสุด', 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่และเวลาที่เพิ่มข้อมูล',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'วันที่และเวลาที่แก้ไขข้อมูลล่าสุด',
);

--------------------------------------- SYSTEM DEFULT TABLE ---------------------------------------

CREATE TABLE sto_attachment ( 
    attachment_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'รหัสไฟล์แนบ (Primary Key)',
    file_name VARCHAR(255) NOT NULL COMMENT 'ชื่อไฟล์ที่ถูกอัปโหลด',
    file_path VARCHAR(255) NOT NULL COMMENT 'เส้นทางที่เก็บไฟล์บนเซิร์ฟเวอร์',
    file_size BIGINT NOT NULL COMMENT 'ขนาดของไฟล์ (ในไบต์)',
    file_type VARCHAR(50) NOT NULL COMMENT 'ประเภทของไฟล์ (เช่น PDF, JPG, PNG)',
    sto_id INT NULL COMMENT 'รหัสรายการที่เชื่อมโยงกับไฟล์แนบ จากตารางที่อัปโหลด [any table.id]',
    sto_tbname VARCHAR(255) NOT NULL COMMENT 'ชื่อตารางที่อัปโหลด',
    sto_field VARCHAR(255) NOT NULL COMMENT 'ชื่อฟิลที่อัปโหลด',
    is_active BOOLEAN DEFAULT TRUE COMMENT 'สถานะการใช้งาน (TRUE = ใช้งาน, FALSE = ไม่ใช้งาน)',
    created_by VARCHAR(100) NULL COMMENT 'ผู้สร้างข้อมูล',
    updated_by VARCHAR(100) NULL COMMENT 'ผู้แก้ไขข้อมูลล่าสุด',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่และเวลาที่เพิ่มข้อมูล',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'วันที่และเวลาที่แก้ไขข้อมูลล่าสุด'
)

CREATE TABLE mst_params_hdr (
    param_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'รหัสพารามิเตอร์ (Primary Key)',
    paramkey_txt VARCHAR(15) NULL COMMENT 'รหัสตัวอักษรเฉพาะของพารามิเตอร์',
    param_code VARCHAR(15) NULL UNIQUE COMMENT 'รหัสเฉพาะของพารามิเตอร์',
    param_nameth VARCHAR(100) NULL COMMENT 'ชื่อพารามิเตอร์ภาษาไทย',
    param_nameeng VARCHAR(100) NULL COMMENT 'ชื่อพารามิเตอร์ภาษาอังกฤษ',
    param_detail VARCHAR(150) NULL COMMENT 'รายละเอียดเพิ่มเติมของพารามิเตอร์',
    param_order INT NULL COMMENT 'ลำดับของพารามิเตอร์',
    is_active BOOLEAN DEFAULT TRUE COMMENT 'สถานะการใช้งาน (TRUE = ใช้งาน, FALSE = ไม่ใช้งาน)',
    created_by VARCHAR(100) NULL COMMENT 'ผู้สร้างข้อมูล',
    updated_by VARCHAR(100) NULL COMMENT 'ผู้แก้ไขข้อมูลล่าสุด',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่และเวลาที่เพิ่มข้อมูล',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'วันที่และเวลาที่แก้ไขข้อมูลล่าสุด'
)

CREATE TABLE mst_params_dtl ( 
    param_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'รหัสพารามิเตอร์ (Primary Key)',
    paramkey_id INT NULL COMMENT 'รหัสคีย์พารามิเตอร์ (Foreign Key ที่อาจเชื่อมกับตารางอื่น) [hardcode id]',
    paramparent_id INT NULL COMMENT 'รหัสพารามิเตอร์พาเรนต์ (สำหรับสร้างโครงสร้างแบบลำดับชั้น) [mst_params_dtl]',
    paramkey_txt VARCHAR(15) NULL COMMENT 'รหัสตัวอักษรเฉพาะของพารามิเตอร์',
    param_code VARCHAR(15) NULL UNIQUE COMMENT 'รหัสเฉพาะของพารามิเตอร์',
    param_nameth VARCHAR(100) NULL COMMENT 'ชื่อพารามิเตอร์ภาษาไทย',
    param_nameeng VARCHAR(100) NULL COMMENT 'ชื่อพารามิเตอร์ภาษาอังกฤษ',
    param_detail VARCHAR(150) NULL COMMENT 'รายละเอียดเพิ่มเติมของพารามิเตอร์',
    param_order INT NULL COMMENT 'ลำดับของพารามิเตอร์',
    paramhdr_id INT NULL COMMENT 'รหัสพารามิเตอร์ส่วนหัว (สัมพันธ์กับตารางหัวข้อ) [mst_params_hdr]',
    is_active BOOLEAN DEFAULT TRUE COMMENT 'สถานะการใช้งาน (TRUE = ใช้งาน, FALSE = ไม่ใช้งาน)',
    created_by VARCHAR(100) NULL COMMENT 'ผู้สร้างข้อมูล',
    updated_by VARCHAR(100) NULL COMMENT 'ผู้แก้ไขข้อมูลล่าสุด',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'วันที่และเวลาที่เพิ่มข้อมูล',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'วันที่และเวลาที่แก้ไขข้อมูลล่าสุด'
);








