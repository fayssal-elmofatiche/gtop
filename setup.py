from setuptools import setup, find_packages

setup(
    name='gittop',
    version='0.1.0',
    packages=find_packages(),
    install_requires=[
        'distro',
        'psutil',
        'screeninfo',
        'rich',
        'pandas',
    ],
    entry_points={
        'console_scripts': [
            'gittop=gittop.main:main',
        ],
    },
    author='Your Name',
    author_email='your.email@example.com',
    description='Git repository and system information display tool with commit heatmap visualization.',
    long_description=open('README.md').read(),
    long_description_content_type='text/markdown',
    url='https://github.com/yourusername/gittop',
    classifiers=[
        'Programming Language :: Python :: 3',
        'License :: OSI Approved :: MIT License',
        'Operating System :: OS Independent',
    ],
    python_requires='>=3.6',
)
